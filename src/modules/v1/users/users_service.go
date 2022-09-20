package users

import (
	"github.com/wildanfaz/backendgolang2_week8/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week8/src/interfaces"
)

type users_service struct {
	repo interfaces.UsersRepo
}

func NewService(repo interfaces.UsersRepo) *users_service {
	return &users_service{repo}
}

func (re *users_service) GetAllUsers() (*models.Users, error) {
	data, err := re.repo.FindAllUsers()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (re *users_service) AddUser(data *models.User) (*models.User, error) {
	data, err := re.repo.SaveUser(data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (re *users_service) UpdateUser(vars string, data *models.User) (*models.User, error) {
	data, err := re.repo.ChangeUser(vars, data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (re *users_service) DeleteUser(vars string, data *models.User) (*models.User, error) {
	data, err := re.repo.RemoveUser(vars, data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// func (re *users_service) SearchUser(r *http.Request) (*models.Users, error) {
// 	data, err := re.repo.FindUser(r)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }
