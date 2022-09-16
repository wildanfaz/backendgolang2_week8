package users

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wildanfaz/backendgolang2_week8/src/database/orm/models"
	"gorm.io/gorm"
)

type users_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *users_repo {
	return &users_repo{db}
}

func (re *users_repo) FindAllUsers() (*models.Users, error) {
	var data models.Users

	result := re.db.Order("created_at desc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed get users")
	}

	return &data, nil
}

func (re *users_repo) SaveUser(data *models.User) (*models.User, error) {
	result := re.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("failed save data")
	}

	return data, nil
}

func (re *users_repo) ChangeUser(r *http.Request, data *models.User) (*models.User, error) {
	vars := mux.Vars(r)

	result := re.db.Model(&data).Where("name = ?", vars["name"]).Updates(data)

	if result.Error != nil {
		return nil, errors.New("failed update data")
	}

	return data, nil
}

func (re *users_repo) RemoveUser(r *http.Request, data *models.User) (*models.User, error) {
	vars := mux.Vars(r)

	result := re.db.Where("name = ?", vars["name"]).Delete(data)

	if result.Error != nil {
		return nil, errors.New("failed delete data")
	}

	return data, nil
}

// func (re *users_repo) FindUser(r *http.Request) (*models.Users, error) {
// 	var data models.Users

// 	search := r.URL.Query().Get("name")
// 	s := "%" + search + "%"
// 	s = strings.ToLower(s)
// 	result := re.db.Where("LOWER(name) LIKE ?", s).Order("created_at desc").Find(&data)

// 	if result.Error != nil {
// 		return nil, errors.New("failed get users")
// 	}

// 	return &data, nil
// }
