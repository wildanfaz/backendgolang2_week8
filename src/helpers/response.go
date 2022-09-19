package helpers

import (
	"encoding/json"
	"net/http"
)

func Response(data interface{}, showData bool, w http.ResponseWriter, status int, msg string, err error) {
	var result = make(map[string]interface{})
	var desc string

	switch status {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
		w.WriteHeader(201)
	case 304:
		desc = "Not Modified"
		http.Error(w, "", http.StatusNotModified)
	case 400:
		desc = "Bad Request"
		http.Error(w, "", http.StatusBadRequest)
	case 401:
		desc = "Unauthorized"
		http.Error(w, "", http.StatusUnauthorized)
	case 404:
		desc = "Unauthorized"
		http.Error(w, "", http.StatusNotFound)
	case 500:
		desc = "Internal Server Error"
		http.Error(w, "", http.StatusInternalServerError)
	case 501:
		desc = "Bad Gateway"
		http.Error(w, "", http.StatusBadGateway)
	default:
		desc = ""
	}

	if err != nil {
		result["http code"] = status
		result["http description"] = desc
		result["msg"] = msg
		result["error"] = err.Error()
	} else if showData == true {
		result["http code"] = status
		result["http description"] = desc
		result["error"] = err
		result["msg"] = msg
		result["result"] = data
	} else {
		result["http code"] = status
		result["http description"] = desc
		result["error"] = err
		result["msg"] = msg
	}
	json.NewEncoder(w).Encode(result)
}
