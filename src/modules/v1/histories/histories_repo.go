package histories

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wildanfaz/backendgolang2_week8/src/database/orm/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		return db.Select("vehicle_id, vehicle_name, created_at, updated_at")
	}).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("user_id, name, email, created_at, updated_at")
	}).Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed get users")
	}

	return &data, nil
}

func (re *histories_repo) SaveHistory(data *models.History) (*models.History, error) {
	result := re.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("failed save history")
	}

	return data, nil
}

func (re *histories_repo) ChangeHistory(r *http.Request, data *models.History) (*models.History, error) {
	vars := mux.Vars(r)

	var check int64

	re.db.Model(&data).Where("history_id = ?", vars["history_id"]).Count(&check)
	checkName := check > 0

	if checkName == false {
		return nil, errors.New("history is not exists")
	}

	result := re.db.Model(&data).Where("history_id = ?", vars["history_id"]).Updates(data)

	if result.Error != nil {
		return nil, errors.New("failed update history")
	}

	return data, nil
}

func (re *histories_repo) RemoveHistory(r *http.Request, data *models.History) (*models.History, error) {
	vars := mux.Vars(r)

	var check int64

	re.db.Model(&data).Where("history_id = ?", vars["history_id"]).Count(&check)
	checkName := check > 0

	if checkName == false {
		return nil, errors.New("history is not exists")
	}

	result := re.db.Delete(data, vars["history_id"])

	if result.Error != nil {
		return nil, errors.New("failed delete history")
	}

	return data, nil
}

func (re *histories_repo) FindHistory(r *http.Request) (*models.Histories, error) {
	var data models.Histories

	search := r.URL.Query().Get("vehicle_id")
	result := re.db.Preload("Vehicle").Preload("User").Where("vehicle_id = ?", search).Order("created_at desc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed get users")
	}

	return &data, nil
}

func (re *histories_repo) FindVehicles() (*models.Histories, error) {
	var data models.Histories

	re.db.Preload(clause.Associations).Select("vehicle_id")

	return &data, nil
}
