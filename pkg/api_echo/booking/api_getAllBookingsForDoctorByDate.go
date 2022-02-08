package booking

import (
	"blockmedic/pkg/env"
	_ "blockmedic/pkg/env"
	"blockmedic/pkg/model"
	op "blockmedic/pkg/operation/booking"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"time"
)

const (
	Time = "02-01-2006"
)

func GetAllBookingsForDoctorByDate(c echo.Context) ([]*model.Patient_Booking, error) {
	date := c.QueryParam("date")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)
	bday, _ := time.Parse(Time, date)
	result, err := op.BookingsForDoctorByDate(&bday, claims.Sub)
	return result, err
}
