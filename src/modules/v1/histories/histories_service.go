package histories

import (
	"net/http"

	"github.com/wildanfaz/backendgolang2_week8/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week8/src/interfaces"
)

type histories_service struct {
	repo interfaces.HistoriesRepo
}

func NewService(repo interfaces.HistoriesRepo) *histories_service {
	return &histories_service{repo}
}

func (re *histories_service) GetAllHistories() (*models.Histories, error) {
	data, err := re.repo.FindAllHistories()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (re *histories_service) AddHistory(data *models.History) (*models.History, error) {
	data, err := re.repo.SaveHistory(data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (re *histories_service) UpdateHistory(r *http.Request, data *models.History) (*models.History, error) {
	data, err := re.repo.ChangeHistory(r, data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (re *histories_service) DeleteHistory(r *http.Request, data *models.History) (*models.History, error) {
	data, err := re.repo.RemoveHistory(r, data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (re *histories_service) SearchHistory(r *http.Request) (*models.Histories, error) {
	data, err := re.repo.FindHistory(r)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (re *histories_service) PopularVehicles() (*models.Histories, error) {
	data, err := re.repo.FindVehicles()

	if err != nil {
		return nil, err
	}

	return data, nil
}
