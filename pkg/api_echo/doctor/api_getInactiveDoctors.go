package doctor

import (
	"blockmedic/pkg/model"
	op "blockmedic/pkg/operation/doctor"
	"github.com/labstack/echo/v4"
)

func GetNewDoctors(c echo.Context) ([]*model.Doctor, error) {
	result, err := op.GetNewDoctors()
	return result, err
}
