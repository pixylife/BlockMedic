package doctor

import (
	"blockmedic/pkg/env"
	"blockmedic/pkg/model"
)

func GetDoctor(id int) (*model.Doctor, error) {
	db := env.RDB
	doctor := model.Doctor{}
	err := doctor.PreloadDoctor(db).First(&doctor, id).Error
	summary := RateSummary{}
	db.Model(model.Rating{}).Select("AVG(rate) AS rate").Where("doctor_id = ?", id).Scan(&summary)
	doctor.Rate = summary.Rate
	if err != nil {
		return nil, err
	} else {
		return &doctor, nil
	}
}
