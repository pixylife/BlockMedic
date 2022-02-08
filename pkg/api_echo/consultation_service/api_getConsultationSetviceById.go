package consultation_service

import (
	_ "blockmedic/pkg/env"
	"blockmedic/pkg/model"
	op "blockmedic/pkg/operation/consultation_service"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetConsultationServicesById(c echo.Context) (*model.ConsultationService, error) {
	id := c.Param("id")
	consultationServiceId, _ := strconv.ParseInt(id, 10, 64)
	result, err := op.ConsultationServiceById(consultationServiceId)
	return result, err
}
