package histories

import (
	"github.com/wildanfaz/backendgolang2_week8/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week8/src/interfaces"
)

type histories_service struct {
	repo interfaces.HistoriesRepo
}

func NewService(repo interfaces.HistoriesRepo) *histories_service {
	return &histories_service{repo}
}

func (svc *histories_service) GetAllHistories() (*models.Histories, error) {
	data, err := svc.repo.FindAllHistories()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (svc *histories_service) AddHistory(data *models.History) (*models.History, error) {
	data, err := svc.repo.SaveHistory(data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (svc *histories_service) UpdateHistory(vars string, data *models.History) (*models.History, error) {
	data, err := svc.repo.ChangeHistory(vars, data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (svc *histories_service) DeleteHistory(vars string, data *models.History) (*models.History, error) {
	data, err := svc.repo.RemoveHistory(vars, data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (svc *histories_service) SearchHistory(search string) (*models.Histories, error) {
	data, err := svc.repo.FindHistory(search)

	if err != nil {
		return nil, err
	}

	return data, nil
}
