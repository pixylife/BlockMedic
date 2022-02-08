package hospital

import (
	"blockmedic/pkg/model"
	op "blockmedic/pkg/operation/hospital"
	"github.com/labstack/echo/v4"
)

func GetHospitals(c echo.Context) ([]*model.Hospital, error) {
	result, err := op.GetHospitals()
	return result, err
}
