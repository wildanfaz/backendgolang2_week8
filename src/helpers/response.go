package helpers

import (
	"encoding/json"
	"net/http"
)

func Response(data interface{}, w http.ResponseWriter, status int, msg string, err error) {
	var result = make(map[string]interface{})
	var desc string

	switch status {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 304:
		desc = "Not Modified"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unauthorized"
	case 500:
		desc = "Internal Server Error"
	case 501:
		desc = "Bad Gateway"
	default:
		desc = ""
	}

	if err != nil {
		result["http code status"] = status
		result["http description"] = desc
		result["error"] = err.Error()
		result["msg"] = msg
		result["result"] = data
	} else {
		result["http status"] = status
		result["response description"] = desc
		result["error"] = err
		result["msg"] = msg
		result["result"] = data
	}

	json.NewEncoder(w).Encode(result)
}
