package interfaces

import (
	"net/http"

	"github.com/wildanfaz/backendgolang2_week8/src/database/orm/models"
)

type HistoriesRepo interface {
	FindAllHistories() (*models.Histories, error)
	SaveHistory(data *models.History) (*models.History, error)
	ChangeHistory(r *http.Request, data *models.History) (*models.History, error)
	RemoveHistory(r *http.Request, data *models.History) (*models.History, error)
	FindHistory(r *http.Request) (*models.Histories, error)
}

type HistoriesService interface {
	GetAllHistories() (*models.Histories, error)
	AddHistory(data *models.History) (*models.History, error)
	UpdateHistory(r *http.Request, data *models.History) (*models.History, error)
	DeleteHistory(r *http.Request, data *models.History) (*models.History, error)
	SearchHistory(r *http.Request) (*models.Histories, error)
}
