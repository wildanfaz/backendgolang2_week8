package users

import (
	"encoding/json"
	"net/http"

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
		helpers.Response(data, w, 400, "", "GET", err)
	} else {
		helpers.Response(data, w, 201, "success get data", "GET", nil)
	}
}

func (re *users_ctrl) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.User

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(datas, w, 400, "", "POST", err)
	} else {
		data, err := re.svc.AddUser(&datas)
		if err != nil {
			helpers.Response(data, w, 400, "", "POST", err)
		} else {
			helpers.Response(data, w, 201, "success add data", "POST", nil)
		}
	}
}

func (re *users_ctrl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.User

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.Response(datas, w, 400, "", "PUT", err)
	} else {

	}

	data, err := re.svc.UpdateUser(r, &datas)
	if err != nil {
		helpers.Response(datas, w, 400, "", "PUT", err)
	} else {
		helpers.Response(data, w, 200, "success update data", "PUT", nil)
	}
}

func (re *users_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.User

	data, err := re.svc.DeleteUser(r, &datas)
	if err != nil {
		helpers.Response(data, w, 400, "", "DELETE", err)
	} else {
		helpers.Response(data, w, 200, "success update data", "PUT", nil)
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
