package patient

import (
	"blockmedic/pkg/model"
	op "blockmedic/pkg/operation/patient"
	"github.com/labstack/echo/v4"
)

func GetPatients(c echo.Context) ([]*model.Patient, error) {
	result, err := op.GetPatients()
	return result, err

}
