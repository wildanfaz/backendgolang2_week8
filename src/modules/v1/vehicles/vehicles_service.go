package vehicles

import (
	"github.com/wildanfaz/backendgolang2_week8/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week8/src/interfaces"
)

type vehicles_service struct {
	repo interfaces.VehiclesRepo
}

func NewService(repo interfaces.VehiclesRepo) *vehicles_service {
	return &vehicles_service{repo}
}

func (svc *vehicles_service) GetAllVehicles() (*models.Vehicles, error) {
	data, err := svc.repo.FindAllVehicles()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (svc *vehicles_service) AddVehicle(data *models.Vehicle) (*models.Vehicle, error) {
	data, err := svc.repo.SaveVehicle(data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (svc *vehicles_service) UpdateVehicle(vars string, data *models.Vehicle) (*models.Vehicle, error) {
	data, err := svc.repo.ChangeVehicle(vars, data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (svc *vehicles_service) DeleteVehicle(vars string, data *models.Vehicle) (*models.Vehicle, error) {
	data, err := svc.repo.RemoveVehicle(vars, data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (svc *vehicles_service) SearchVehicle(search string) (*models.Vehicles, error) {
	data, err := svc.repo.FindVehicle(search)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (svc *vehicles_service) PopularVehicles() (*models.Vehicles, error) {
	data, err := svc.repo.RatingVehicles()

	if err != nil {
		return nil, err
	}

	return data, nil
}
