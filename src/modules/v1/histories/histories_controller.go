package histories

import (
	"encoding/json"
	"net/http"

	"github.com/wildanfaz/backendgolang2_week8/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week8/src/helpers"
	"github.com/wildanfaz/backendgolang2_week8/src/interfaces"
)

type histories_ctrl struct {
	svc interfaces.HistoriesService
}

func NewCtrl(svc interfaces.HistoriesService) *histories_ctrl {
	return &histories_ctrl{svc}
}

func (re *histories_ctrl) GetAllHistories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := re.svc.GetAllHistories()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	helpers.Response(data, w, 200, "success get data", nil)
}

func (re *histories_ctrl) AddHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.History

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	data, err := re.svc.AddHistory(&datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(data)
}

func (re *histories_ctrl) UpdateHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.History

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	data, err := re.svc.UpdateHistory(r, &datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(data)
}

func (re *histories_ctrl) DeleteHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.History

	data, err := re.svc.DeleteHistory(r, &datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(data)
}

func (re *histories_ctrl) SearchHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := re.svc.SearchHistory(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(data)
}
