package consultation_service

import (
	"blockmedic/pkg/env"
	_ "blockmedic/pkg/env"
	"blockmedic/pkg/model"
	op "blockmedic/pkg/operation/consultation_service"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func UpdateConsultationService(c echo.Context) (*model.ConsultationService, error) {
	consultationService := model.ConsultationService{}
	if error := c.Bind(&consultationService); error != nil {
		return nil, error
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)

	result, err := op.ConsultationServiceUpdate(&consultationService, claims.Sub)
	return result, err
}
