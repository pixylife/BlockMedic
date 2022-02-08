package consultation_service

import (
	_ "blockmedic/pkg/env"
	"blockmedic/pkg/model"
	op "blockmedic/pkg/operation/consultation_service"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func GetAllConsultationServices(c echo.Context) ([]*model.ConsultationService, error) {
	result, err := op.ConsultationServicesGetAll()
	return result, err
}
