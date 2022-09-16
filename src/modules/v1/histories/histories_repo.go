package histories

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
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

	result := re.db.Preload("Vehicle").Preload("User").Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed get histories")
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

	result := re.db.Model(&data).Where("history_id = ?", vars["history_id"]).Updates(data)

	if result.Error != nil {
		return nil, errors.New("failed update history")
	}

	return data, nil
}

func (re *histories_repo) RemoveHistory(r *http.Request, data *models.History) (*models.History, error) {
	vars := mux.Vars(r)

	result := re.db.Delete(data, vars["history_id"])

	if result.Error != nil {
		return nil, errors.New("failed delete history")
	}

	return data, nil
}

func (re *histories_repo) FindHistory(r *http.Request) (*models.Histories, error) {
	var data models.Histories

	search := r.URL.Query().Get("vehicle_id")
	result := re.db.Preload("Vehicle").Preload("User").Where("vehicle_id = ?", search).Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed get users")
	}

	return &data, nil
}
