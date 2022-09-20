package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wildanfaz/backendgolang2_week8/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week8/src/helpers"
	"github.com/wildanfaz/backendgolang2_week8/src/interfaces"
)

type users_ctrl struct {
	svc interfaces.UsersService
}

func NewCtrl(svc interfaces.UsersService) *users_ctrl {
	return &users_ctrl{svc}
}

func (re *users_ctrl) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := re.svc.GetAllUsers()

	if err != nil {
		helpers.Response(data, true, w, 400, "failed get data", err)
	} else {
		helpers.Response(data, true, w, 200, "success get data", nil)
	}
}

func (re *users_ctrl) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.User

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(datas, false, w, 400, "failed add data", err)
	} else {
		data, err := re.svc.AddUser(&datas)

		if err != nil {
			helpers.Response(data, false, w, 400, "failed add data", err)
		} else {
			helpers.Response(data, false, w, 201, "success add data", nil)
		}
	}
}

func (re *users_ctrl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.User

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(datas, false, w, 400, "failed update data", err)
	} else {
		vars := mux.Vars(r)
		data, err := re.svc.UpdateUser(vars["name"], &datas)

		if err != nil {
			helpers.Response(datas, false, w, 400, "failed update data", err)
		} else {
			helpers.Response(data, false, w, 200, "success update data", nil)
		}
	}
}

func (re *users_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.User

	vars := mux.Vars(r)

	data, err := re.svc.DeleteUser(vars["name"], &datas)

	if err != nil {
		helpers.Response(data, false, w, 400, "failed delete data", err)
	} else {
		helpers.Response(data, false, w, 200, "success delete data", nil)
	}
}

// func (re *users_ctrl) SearchUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	data, err := re.svc.SearchUser(r)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	}

// 	json.NewEncoder(w).Encode(data)
// }
