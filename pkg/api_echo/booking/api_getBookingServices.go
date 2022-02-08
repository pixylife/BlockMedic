package booking

import (
	"blockmedic/pkg/model"
	op "blockmedic/pkg/operation/booking"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetAllBookingsForDoctorById(c echo.Context) ([]*model.BookingService, error) {
	id := c.Param("id")
	doctorID, _ := strconv.ParseInt(id, 10, 64)
	result, err := op.BookingRequestsForDoctor(doctorID)
	return result, err
}
