package availability

import (
	"blockmedic/pkg/env"
	"blockmedic/pkg/model"
	op "blockmedic/pkg/operation/availability"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func GetAvailabilityByDate(c echo.Context) ([]*model.Availability, error) {
	date := c.QueryParam("date")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)
	result, err := op.GetAvailabilityByDate(date, claims.Sub)
	return result, err
}
