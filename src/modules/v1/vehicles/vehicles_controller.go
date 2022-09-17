package vehicles

import (
	"encoding/json"
	"net/http"

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

func (re *vehicles_ctrl) GetAllVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := re.svc.GetAllVehicles()

	if err != nil {
		helpers.Response(data, w, 400, "", "GET", err)
	} else {
		helpers.Response(data, w, 200, "success get data", "GET", nil)
	}
}

func (re *vehicles_ctrl) AddVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.Vehicle

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(datas, w, 400, "", "POST", err)
	} else {
		data, err := re.svc.AddVehicle(&datas)
		if err != nil {
			helpers.Response(data, w, 400, "", "POST", err)
		} else {
			helpers.Response(data, w, 200, "success add data", "POST", nil)
		}
	}
}

func (re *vehicles_ctrl) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.Vehicle

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(datas, w, 400, "", "PUT", err)
	} else {
		data, err := re.svc.UpdateVehicle(r, &datas)
		if err != nil {
			helpers.Response(data, w, 400, "", "PUT", err)
		} else {
			helpers.Response(data, w, 200, "success update data", "PUT", nil)
		}
	}
}

func (re *vehicles_ctrl) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.Vehicle

	data, err := re.svc.DeleteVehicle(r, &datas)
	if err != nil {
		helpers.Response(data, w, 400, "", "DELETE", err)
	} else {
		helpers.Response(data, w, 200, "success delete data", "DELETE", nil)
	}
}

func (re *vehicles_ctrl) SearchVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := re.svc.SearchVehicle(r)

	if err != nil {
		helpers.Response(data, w, 400, "", "GET", err)
	} else {
		helpers.Response(data, w, 200, "success search data", "GET", nil)
	}
}
