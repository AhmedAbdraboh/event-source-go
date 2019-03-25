package base

import (
	"encoding/json"
	"net/http"
	"roosh-app/helpers"
)

type Controller struct {

}

func (controller *Controller) RespondWithJson(w http.ResponseWriter, payload interface{}, statusCode int)  {
	jsonResponse, err := json.Marshal(payload)
	helpers.CheckPanicError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(jsonResponse)
	helpers.CheckPanicError(err)
}

