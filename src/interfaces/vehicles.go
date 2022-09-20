package interfaces

import (
	"github.com/wildanfaz/backendgolang2_week8/src/database/orm/models"
)

type VehiclesRepo interface {
	FindAllVehicles() (*models.Vehicles, error)
	SaveVehicle(data *models.Vehicle) (*models.Vehicle, error)
	ChangeVehicle(vars string, data *models.Vehicle) (*models.Vehicle, error)
	RemoveVehicle(vars string, data *models.Vehicle) (*models.Vehicle, error)
	FindVehicle(search string) (*models.Vehicles, error)
	RatingVehicles() (*models.Vehicles, error)
}

type VehiclesService interface {
	GetAllVehicles() (*models.Vehicles, error)
	AddVehicle(data *models.Vehicle) (*models.Vehicle, error)
	UpdateVehicle(vars string, data *models.Vehicle) (*models.Vehicle, error)
	DeleteVehicle(vars string, data *models.Vehicle) (*models.Vehicle, error)
	SearchVehicle(search string) (*models.Vehicles, error)
	PopularVehicles() (*models.Vehicles, error)
}
