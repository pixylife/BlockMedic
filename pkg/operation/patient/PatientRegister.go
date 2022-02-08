package patient

import (
	"blockmedic/pkg/env"
	"blockmedic/pkg/model"
)

func PatientRegister(patient *model.Patient) (*model.Patient, error) {
	db := env.RDB
	err := db.Create(patient).Error
	if err != nil {
		return nil, err
	} else {
		return patient, nil
	}
}
