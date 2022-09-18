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
		helpers.Response(data, w, 400, "failed get data", "GET", err)
	} else {
		helpers.Response(data, w, 200, "success get data", "GET", nil)
	}
}

func (re *histories_ctrl) AddHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.History

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(datas, w, 400, "failed add data", "POST", err)
	} else {
		data, err := re.svc.AddHistory(&datas)

		if err != nil {
			helpers.Response(data, w, 400, "failed add data", "POST", err)
		} else {
			helpers.Response(data, w, 201, "success add data", "POST", nil)
		}
	}
}

func (re *histories_ctrl) UpdateHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.History

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(datas, w, 400, "failed update data", "PUT", err)
	} else {
		data, err := re.svc.UpdateHistory(r, &datas)

		if err != nil {
			helpers.Response(data, w, 400, "failed update data", "PUT", err)
		} else {
			helpers.Response(data, w, 200, "success update data", "PUT", nil)
		}
	}
}

func (re *histories_ctrl) DeleteHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.History

	data, err := re.svc.DeleteHistory(r, &datas)

	if err != nil {
		helpers.Response(data, w, 400, "failed delete data", "DELETE", err)
	} else {
		helpers.Response(data, w, 200, "success delete data", "DELETE", nil)
	}
}

func (re *histories_ctrl) SearchHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := re.svc.SearchHistory(r)

	if err != nil {
		helpers.Response(data, w, 400, "failed search data", "GET", err)
	} else {
		helpers.Response(data, w, 200, "success search data", "GET", nil)
	}
}
