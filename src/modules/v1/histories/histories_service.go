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

func (re *histories_service) UpdateHistory(vars string, data *models.History) (*models.History, error) {
	data, err := re.repo.ChangeHistory(vars, data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (re *histories_service) DeleteHistory(vars string, data *models.History) (*models.History, error) {
	data, err := re.repo.RemoveHistory(vars, data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (re *histories_service) SearchHistory(search string) (*models.Histories, error) {
	data, err := re.repo.FindHistory(search)

	if err != nil {
		return nil, err
	}

	return data, nil
}
