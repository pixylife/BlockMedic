package consultation_service

import (
	"blockmedic/pkg/env"
	"blockmedic/pkg/model"
)

func ConsultationServiceById(id int64) (*model.ConsultationService, error) {
	db := env.RDB
	consultationService := model.ConsultationService{}
	err := consultationService.PreloadConsultationService(db).Model(model.ConsultationService{}).First(&consultationService, id).Error
	if err != nil {
		return nil, err
	} else {
		return &consultationService, nil
	}
}
