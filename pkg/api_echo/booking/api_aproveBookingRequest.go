package booking

import (
	"blockmedic/pkg/env"
	_ "blockmedic/pkg/env"
	"blockmedic/pkg/model"
	op "blockmedic/pkg/operation/booking"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strconv"
)

func ApproveBookingRequest(c echo.Context) (*model.BookingRequest, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	approve, _ := strconv.ParseBool(c.QueryParam("approve"))
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)

	status := ""
	if approve {
		status = env.STATUS_APPROVED
	} else {
		status = env.STATUS_REJECTED
	}

	result, err := op.ApproveBookingRequest(int64(id), status, claims.Sub)
	return result, err
}
