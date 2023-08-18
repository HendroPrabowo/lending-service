package account

import (
	"encoding/json"
	"net/http"

	"lending-service/utility/response"
)

type controller struct {
	service service
}

func newController(service service) controller {
	return controller{
		service: service,
	}
}

func (c controller) Register(w http.ResponseWriter, r *http.Request) {
	var dto AccountDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.ErrorWithMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.service.ProcessRegister(dto); err != nil {
		response.ErrorWrapped(w, err)
		return
	}

	response.OkWithMessage(w, "success")
}
