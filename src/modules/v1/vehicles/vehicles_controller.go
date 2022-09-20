package vehicles

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wildanfaz/backendgolang2_week8/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week8/src/helpers"
	"github.com/wildanfaz/backendgolang2_week8/src/interfaces"
)

type vehicles_ctrl struct {
	svc interfaces.VehiclesService
}

func NewCtrl(svc interfaces.VehiclesService) *vehicles_ctrl {
	return &vehicles_ctrl{svc}
}

func (ctrl *vehicles_ctrl) GetAllVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := ctrl.svc.GetAllVehicles()

	if err != nil {
		helpers.Response(data, true, w, 400, "failed get data", err)
	} else {
		helpers.Response(data, true, w, 200, "success get data", nil)
	}
}

func (ctrl *vehicles_ctrl) AddVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.Vehicle

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(datas, false, w, 400, "failed add data", err)
	} else {
		data, err := ctrl.svc.AddVehicle(&datas)

		if err != nil {
			helpers.Response(data, false, w, 400, "failed add data", err)
		} else {
			helpers.Response(data, false, w, 201, "success add data", nil)
		}
	}
}

func (ctrl *vehicles_ctrl) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.Vehicle

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(datas, false, w, 400, "failed update data", err)
	} else {
		vars := mux.Vars(r)
		data, err := ctrl.svc.UpdateVehicle(vars["vehicle_id"], &datas)

		if err != nil {
			helpers.Response(data, false, w, 400, "failed update data", err)
		} else {
			helpers.Response(data, false, w, 200, "success update data", nil)
		}
	}
}

func (ctrl *vehicles_ctrl) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.Vehicle

	vars := mux.Vars(r)
	data, err := ctrl.svc.DeleteVehicle(vars["vehicle_id"], &datas)

	if err != nil {
		helpers.Response(data, false, w, 400, "failed delete data", err)
	} else {
		helpers.Response(data, false, w, 200, "success delete data", nil)
	}
}

func (ctrl *vehicles_ctrl) SearchVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	search := r.URL.Query().Get("vehicle_name")
	data, err := ctrl.svc.SearchVehicle(search)

	if err != nil {
		helpers.Response(data, true, w, 400, "failed search data", err)
	} else {
		helpers.Response(data, true, w, 200, "success search data", nil)
	}
}

func (ctrl *vehicles_ctrl) PopularVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := ctrl.svc.PopularVehicles()

	if err != nil {
		helpers.Response(data, true, w, 400, "failed get data", err)
	} else {
		helpers.Response(data, true, w, 200, "success get data", nil)
	}
}
