package histories

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

func (ctrl *histories_ctrl) GetAllHistories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := ctrl.svc.GetAllHistories()

	if err != nil {
		helpers.Response(data, true, w, 400, "failed get data", err)
	} else {
		helpers.Response(data, true, w, 200, "success get data", nil)
	}
}

func (ctrl *histories_ctrl) AddHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.History

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(datas, false, w, 400, "failed add data", err)
	} else {
		data, err := ctrl.svc.AddHistory(&datas)

		if err != nil {
			helpers.Response(data, false, w, 400, "failed add data", err)
		} else {
			helpers.Response(data, false, w, 201, "success add data", nil)
		}
	}
}

func (ctrl *histories_ctrl) UpdateHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.History

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(datas, false, w, 400, "failed update data", err)
	} else {
		vars := mux.Vars(r)
		data, err := ctrl.svc.UpdateHistory(vars["history_id"], &datas)

		if err != nil {
			helpers.Response(data, false, w, 400, "failed update data", err)
		} else {
			helpers.Response(data, false, w, 200, "success update data", nil)
		}
	}
}

func (ctrl *histories_ctrl) DeleteHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.History

	vars := mux.Vars(r)
	data, err := ctrl.svc.DeleteHistory(vars["history_id"], &datas)

	if err != nil {
		helpers.Response(data, false, w, 400, "failed delete data", err)
	} else {
		helpers.Response(data, false, w, 200, "success delete data", nil)
	}
}

func (ctrl *histories_ctrl) SearchHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	search := r.URL.Query().Get("vehicle_id")
	data, err := ctrl.svc.SearchHistory(search)

	if err != nil {
		helpers.Response(data, true, w, 400, "failed search data", err)
	} else {
		helpers.Response(data, true, w, 200, "success search data", nil)
	}
}
