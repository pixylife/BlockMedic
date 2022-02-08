package availability

import (
	"blockmedic/pkg/env"
	"blockmedic/pkg/model"
)

func GetAvailability(id int64) (*model.Availability, error) {
	db := env.RDB
	availability := model.Availability{}
	availability.PreloadAvailability(db).First(&availability, id)
	return &availability, nil
}
