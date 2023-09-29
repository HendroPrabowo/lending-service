package response

import (
	"encoding/json"
	"net/http"

	"lending-service/utility/wraped_error"

	log "github.com/sirupsen/logrus"
)

func Ok(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	writeResponse(w, resp)
}

func OkWithMessage(w http.ResponseWriter, resp interface{}) {
	mapResp := map[string]interface{}{"message": resp}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	writeResponse(w, mapResp)
}

func Error(w http.ResponseWriter, status int, resp interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	writeResponse(w, resp)
}

func ErrorWithMessage(w http.ResponseWriter, status int, resp interface{}) {
	mapResp := map[string]interface{}{"message": resp}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	writeResponse(w, mapResp)
}

func ErrorWrapped(w http.ResponseWriter, err *wraped_error.Error) {
	if err.StatusCode == http.StatusInternalServerError {
		//bugsnag.Notify(errors.New(err.Err, 1))
	}

	mapResp := map[string]interface{}{"message": err.Err.Error()}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(err.StatusCode)
	writeResponse(w, mapResp)
}

func writeResponse(w http.ResponseWriter, resp interface{}) {
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Error(err)
	}
	w.Write(jsonResp)
}
