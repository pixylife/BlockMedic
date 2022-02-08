package consultation_service

import (
	"blockmedic/pkg/env"
	_ "blockmedic/pkg/env"
	"blockmedic/pkg/model"
	op "blockmedic/pkg/operation/consultation_service"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strconv"
)

func DeleteConsultationService(c echo.Context) (*model.ConsultationService, error) {
	id := c.Param("id")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)
	consultationServiceId, _ := strconv.ParseInt(id, 10, 64)
	result, err := op.ConsultationServiceDelete(consultationServiceId, claims.Sub, claims.Auth)
	return result, err
}
