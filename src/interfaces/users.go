package interfaces

import (
	"net/http"

	"github.com/wildanfaz/backendgolang2_week8/src/database/orm/models"
)

type UsersRepo interface {
	FindAllUsers() (*models.Users, error)
	SaveUser(data *models.User) (*models.User, error)
	ChangeUser(r *http.Request, data *models.User) (*models.User, error)
	RemoveUser(r *http.Request, data *models.User) (*models.User, error)
	FindUser(r *http.Request) (*models.Users, error)
}

type UsersService interface {
	GetAllUsers() (*models.Users, error)
	AddUser(data *models.User) (*models.User, error)
	UpdateUser(r *http.Request, data *models.User) (*models.User, error)
	DeleteUser(r *http.Request, data *models.User) (*models.User, error)
	SearchUser(r *http.Request) (*models.Users, error)
}
