package availability

import (
	"blockmedic/pkg/env"
	"blockmedic/pkg/model"
)

func AddAvailabiility(availability *model.Availability, email string) (*model.Availability, error) {
	db := env.RDB
	user := model.User{}
	user.PreloadDoctor(db).Where("email = ?", email).First(&user)
	availability.Doctor = user.Doctor
	availability.IsFree = true
	hospital := model.Hospital{}
	db.Model(model.Hospital{}).First(&hospital, availability.HospitalID)
	availability.Hospital = &hospital
	err := db.Create(availability).Error
	if err != nil {
		return nil, err
	} else {
		return availability, nil
	}
}
