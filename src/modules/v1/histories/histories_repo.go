package histories

import (
	"errors"

	"github.com/wildanfaz/backendgolang2_week8/src/database/orm/models"
	"gorm.io/gorm"
)

type histories_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *histories_repo {
	return &histories_repo{db}
}

func (re *histories_repo) FindAllHistories() (*models.Histories, error) {
	var data models.Histories

	result := re.db.Order("created_at desc").Preload("Vehicle", func(db *gorm.DB) *gorm.DB {
		return db.Select("vehicle_id, vehicle_name, created_at, updated_at, total_rented")
	}).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("user_id, name, email, created_at, updated_at")
	}).Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed get users")
	}

	return &data, nil
}

func (re *histories_repo) SaveHistory(data *models.History) (*models.History, error) {
	var vehicle models.Vehicle

	result := re.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("failed save history")
	}

	re.db.Where("vehicle_id = ?", data.VehicleId).First(&vehicle)

	re.db.Model(&vehicle).Where("vehicle_id = ?", data.VehicleId).Update("total_rented", vehicle.TotalRented+1)

	return data, nil
}

func (re *histories_repo) ChangeHistory(vars string, data *models.History) (*models.History, error) {
	var check int64

	re.db.Model(&data).Where("history_id = ?", vars).Count(&check)
	checkName := check > 0

	if checkName == false {
		return nil, errors.New("history is not exists")
	}

	result := re.db.Model(&data).Where("history_id = ?", vars).Updates(data)

	if result.Error != nil {
		return nil, errors.New("failed update history")
	}

	return data, nil
}

func (re *histories_repo) RemoveHistory(vars string, data *models.History) (*models.History, error) {
	var vehicle models.Vehicle

	var check int64

	re.db.Model(&data).Where("history_id = ?", vars).Count(&check)
	checkName := check > 0

	if checkName == false {
		return nil, errors.New("history is not exists")
	}

	re.db.Where("history_id = ?", vars).First(&data)

	re.db.Where("vehicle_id = ?", data.VehicleId).First(&vehicle)

	re.db.Model(&vehicle).Where("vehicle_id = ?", data.VehicleId).Update("total_rented", vehicle.TotalRented-1)

	result := re.db.Delete(data, vars)

	if result.Error != nil {
		return nil, errors.New("failed delete history")
	}

	return data, nil
}

func (re *histories_repo) FindHistory(search string) (*models.Histories, error) {
	var data models.Histories

	result := re.db.Preload("Vehicle", func(db *gorm.DB) *gorm.DB {
		return db.Select("vehicle_id, vehicle_name, created_at, updated_at, total_rented")
	}).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("user_id, name, email, created_at, updated_at")
	}).Where("vehicle_id = ?", search).Order("created_at desc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed get users")
	}

	return &data, nil
}
