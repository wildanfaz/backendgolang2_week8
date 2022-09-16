package interfaces

import (
	"net/http"

	"github.com/wildanfaz/backendgolang2_week8/src/database/orm/models"
)

type VehiclesRepo interface {
	FindAllVehicles() (*models.Vehicles, error)
	SaveVehicle(data *models.Vehicle) (*models.Vehicle, error)
	ChangeVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error)
	RemoveVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error)
}

type VehiclesService interface {
	GetAllVehicles() (*models.Vehicles, error)
	AddVehicle(data *models.Vehicle) (*models.Vehicle, error)
	UpdateVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error)
	DeleteVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error)
}
