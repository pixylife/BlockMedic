package patient

import (
	"blockmedic/pkg/model"
	op "blockmedic/pkg/operation/patient"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetPatient(c echo.Context) (*model.Patient, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := op.GetPatient(id)
	return result, err

}
