package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
)

func Response(data interface{}, w http.ResponseWriter, status int, msg string, method string, err error) {
	var result = make(map[string]interface{})
	var desc string

	switch status {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 304:
		desc = "Not Modified"
		http.Error(w, msg, http.StatusNotModified)
	case 400:
		desc = "Bad Request"
		http.Error(w, msg, http.StatusBadRequest)
	case 401:
		desc = "Unauthorized"
		http.Error(w, msg, http.StatusUnauthorized)
	case 500:
		desc = "Internal Server Error"
		http.Error(w, msg, http.StatusInternalServerError)
	case 501:
		desc = "Bad Gateway"
		http.Error(w, msg, http.StatusBadGateway)
	default:
		desc = ""
	}

	checkMethod := []string{"GET", "POST", "PUT", "DELETE"}
	for _, v := range checkMethod {
		if method == v {
			checkMethod = append(checkMethod, "add length")
			break
		}
	}

	if len(checkMethod) == 4 {
		err = errors.New("invalid method in controller")
	}

	if err != nil {
		result["http code"] = status
		result["http description"] = desc
		result["error"] = err.Error()
		json.NewEncoder(w).Encode(result)
	} else if method == "GET" {
		result["http code"] = status
		result["http description"] = desc
		result["error"] = err
		result["msg"] = msg
		result["result"] = data
		json.NewEncoder(w).Encode(result)
	} else {
		result["http code"] = status
		result["http description"] = desc
		result["error"] = err
		result["msg"] = msg
		json.NewEncoder(w).Encode(result)
	}
}

// func Error()
