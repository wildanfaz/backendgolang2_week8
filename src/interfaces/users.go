package interfaces

import (
	"github.com/wildanfaz/backendgolang2_week8/src/database/orm/models"
)

type UsersRepo interface {
	FindAllUsers() (*models.Users, error)
	SaveUser(data *models.User) (*models.User, error)
	ChangeUser(vars string, data *models.User) (*models.User, error)
	RemoveUser(vars string, data *models.User) (*models.User, error)
	// FindUser(r *http.Request) (*models.Users, error)
}

type UsersService interface {
	GetAllUsers() (*models.Users, error)
	AddUser(data *models.User) (*models.User, error)
	UpdateUser(vars string, data *models.User) (*models.User, error)
	DeleteUser(vars string, data *models.User) (*models.User, error)
	// SearchUser(r *http.Request) (*models.Users, error)
}
