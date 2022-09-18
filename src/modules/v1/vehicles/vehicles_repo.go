package vehicles

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wildanfaz/backendgolang2_week8/src/database/orm/models"
	"gorm.io/gorm"
)

type vehicles_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *vehicles_repo {
	return &vehicles_repo{db}
}

func (re *vehicles_repo) FindAllVehicles() (*models.Vehicles, error) {
	var data models.Vehicles

	result := re.db.Order("created_at desc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed get vehicles")
	}

	return &data, nil
}

func (re *vehicles_repo) SaveVehicle(data *models.Vehicle) (*models.Vehicle, error) {
	result := re.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("failed save vehicle")
	}

	return data, nil
}

func (re *vehicles_repo) ChangeVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error) {
	vars := mux.Vars(r)

	var check int64

	re.db.Model(&data).Where("vehicle_id = ?", vars["vehicle_id"]).Count(&check)
	checkName := check > 0

	if checkName == false {
		return nil, errors.New("vehicle is not exists")
	}

	result := re.db.Model(&data).Where("vehicle_id = ?", vars["vehicle_id"]).Updates(data)

	if result.Error != nil {
		return nil, errors.New("failed update vehicle")
	}

	return data, nil
}

func (re *vehicles_repo) RemoveVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error) {
	vars := mux.Vars(r)

	var check int64

	re.db.Model(&data).Where("vehicle_id = ?", vars["vehicle_id"]).Count(&check)
	checkName := check > 0

	if checkName == false {
		return nil, errors.New("vehicle is not exists")
	}

	result := re.db.Delete(data, vars["vehicle_id"])

	if result.Error != nil {
		return nil, errors.New("failed delete vehicle")
	}

	return data, nil
}

func (re *vehicles_repo) FindVehicle(r *http.Request) (*models.Vehicles, error) {
	var data models.Vehicles

	search := r.URL.Query().Get("vehicle_name")
	s := "%" + search + "%"
	result := re.db.Where("LOWER(vehicle_name) LIKE ?", s).Order("created_at desc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed get users")
	}

	return &data, nil
}

func (re *vehicles_repo) RatingVehicles() (*models.Vehicles, error) {
	var data models.Vehicles

	result := re.db.Order("rating desc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed get vehicles")
	}

	return &data, nil
}
