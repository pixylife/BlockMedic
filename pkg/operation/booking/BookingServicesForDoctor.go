package booking

import (
	"blockmedic/pkg/env"
	"blockmedic/pkg/model"
)

func BookingRequestsForDoctor(id int64) ([]*model.BookingService, error) {
	db := env.RDB
	bookingService := model.BookingService{}
	bookingServices := []*model.BookingService{}
	bookingService.PreloadBookingService(db).Model(model.BookingService{}).Where("doctor_id = ? ", id).Find(&bookingServices)
	return bookingServices, nil
}
