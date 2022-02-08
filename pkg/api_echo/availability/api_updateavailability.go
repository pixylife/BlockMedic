package availability

import (
	"blockmedic/pkg/env"
	_ "blockmedic/pkg/env"
	"blockmedic/pkg/model"
	op "blockmedic/pkg/operation/availability"
	"context"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func UpdateAvailability(c echo.Context, ctx context.Context) (*model.Availability, error) {
	availability := model.Availability{}
	if err := c.Bind(&availability); err != nil {
		return nil, err
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)

	result, err := op.UpdateAvailabiility(&availability, claims.Sub)
	return result, err
}
