package data

import (
	"errors"
	"orderfaz-1/app/database"
	"orderfaz-1/features/logistic"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type logisticQuery struct {
	db *gorm.DB
}

func NewLogistic(db *gorm.DB) logistic.LogisticDataInterface {
	return &logisticQuery{
		db: db,
	}
}

func (repo *logisticQuery) SerachByQuery(origin_name, dastination_name string) ([]logistic.CoreLogistic, error) {
	var logisticDataGorm []logistic.CoreLogistic

	tx := repo.db.Table("logistics").Where("origin_name LIKE ?", "%"+origin_name+"%").
		Where("dastination_name LIKE ?", "%"+dastination_name+"%").
		Find(&logisticDataGorm)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var logisticDataCore []logistic.CoreLogistic
	for _, value := range logisticDataGorm {
		var logisticCore = logistic.CoreLogistic{
			Uuid:            value.Uuid,
			LogisticName:    value.LogisticName,
			Amount:          value.Amount,
			DastinationName: value.DastinationName,
			OriginName:      value.OriginName,
			Duration:        value.Duration,
		}
		logisticDataCore = append(logisticDataCore, logisticCore)
	}

	return logisticDataCore, nil
}

func (repo *logisticQuery) CreateLogistic(input logistic.CoreLogistic) error {
	uuid := uuid.New().String()
	inputLogistic := database.Logistic{
		Uuid:            uuid,
		LogisticName:    input.LogisticName,
		Amount:          input.Amount,
		DastinationName: input.DastinationName,
		OriginName:      input.OriginName,
		Duration:        input.Duration,
	}

	tx := repo.db.Create(&inputLogistic)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}
