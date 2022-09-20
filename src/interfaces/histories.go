package interfaces

import (
	"github.com/wildanfaz/backendgolang2_week8/src/database/orm/models"
)

type HistoriesRepo interface {
	FindAllHistories() (*models.Histories, error)
	SaveHistory(data *models.History) (*models.History, error)
	ChangeHistory(vars string, data *models.History) (*models.History, error)
	RemoveHistory(vars string, data *models.History) (*models.History, error)
	FindHistory(search string) (*models.Histories, error)
}

type HistoriesService interface {
	GetAllHistories() (*models.Histories, error)
	AddHistory(data *models.History) (*models.History, error)
	UpdateHistory(vars string, data *models.History) (*models.History, error)
	DeleteHistory(vars string, data *models.History) (*models.History, error)
	SearchHistory(search string) (*models.Histories, error)
}
