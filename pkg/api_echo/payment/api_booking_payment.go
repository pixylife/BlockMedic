package payment

import (
	"blockmedic/pkg/env"
	_ "blockmedic/pkg/env"
	"blockmedic/pkg/model"
	op "blockmedic/pkg/operation/payment"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func BookingPayment(c echo.Context) (*model.BookingPayment, error) {
	token := c.QueryParam("token")
	payment := model.BookingPayment{}
	if error := c.Bind(&payment); error != nil {
		return nil, error
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)

	result, err := op.ProceedBookingPayment(&payment, token, claims.Sub)
	return result, err
}
