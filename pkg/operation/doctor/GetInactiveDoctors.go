package doctor

import (
	"blockmedic/pkg/env"
	"blockmedic/pkg/model"
)

func GetNewDoctors() ([]*model.Doctor, error) {
	db := env.RDB
	doctors := []*model.Doctor{}
	db.Model(model.Doctor{}).Where("is_new = ?", env.STATUS_NEW).Scan(&doctors)
	return doctors, nil

}
