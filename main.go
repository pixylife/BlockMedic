package main

import (
	logger "blockmedic/pkg/logger"
	model "blockmedic/pkg/model"
	"crypto/subtle"
	"github.com/dgrijalva/jwt-go"
	gorm "github.com/jinzhu/gorm"
)

import (
	"blockmedic/pkg/env"
)
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
import (
	"flag"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
import (
	"blockmedic/pkg/controller_echo"
)

func main() {

	database, err := gorm.Open("mysql", env.DBuser+":"+env.DBpwd+"@tcp("+env.DBhost+":"+env.DBport+")/"+env.DBdb+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		logger.Log.Error(err)
	}
	env.RDB = database

	model.InitModels(database)
	/*
		folderPath := "patient_images"
		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			_ = os.Mkdir(folderPath, os.ModePerm)
		}

		ctx := context.Background()
		context.WithValue(ctx, "myvalues", folderPath)

		folderPath = "doctor_images"
		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			_ = os.Mkdir(folderPath, os.ModePerm)
		}*/

	CreateDefaultUser()
	RunProxy()
}

func RunProxy() {
	flag.Parse()
	defer glog.Flush()
	run()
}

func run() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	r := e.Group("/")
	jwtConfig := middleware.JWTConfig{
		Claims:     &env.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			jwtClaims := user.Claims.(*env.JwtCustomClaims)
			c.Set("ss", jwtClaims.Sub)
			return next(c)
		}
	})

	r.Use(middleware.JWTWithConfig(jwtConfig))
	controller_echo.APIControllerDoctor(r)
	controller_echo.APIControllerHospital(r)
	controller_echo.APIControllerPatient(r)
	controller_echo.APIControllerPayment(r)
	controller_echo.APIControllerConsultationService(r)
	controller_echo.APIControllerRating(r)
	controller_echo.APIControllerBookingService(r)

	u := e.Group("/")
	u.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte("blockmedic")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("blockmedic@123")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	controller_echo.LoginController(u)
	controller_echo.APIControllerHospitalBase(u)

	e.Logger.Fatal(e.Start(":" + env.RestPort))
}

func CreateDefaultUser() {
	db := env.RDB
	users := []*model.User{}
	db.Find(&users)
	var targetUser *model.User
	if len(users) == 0 {
		targetUser = &model.User{}
		targetUser.Role = env.ADMIN
		targetUser.Email = "admin"
		targetUser.Password = "admin"
		targetUser.Name = "Admin"
		db.Create(&targetUser)
	}
}
