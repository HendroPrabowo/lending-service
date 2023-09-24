package account

import (
	"encoding/json"
	"net/http"

	"lending-service/utility/response"
	"lending-service/utility/wraped_error"
)

type controller struct {
	service Service
}

func newController(service Service) controller {
	return controller{
		service: service,
	}
}

func (c controller) Register(w http.ResponseWriter, r *http.Request) {
	var dto AccountDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.ErrorWrapped(w, wraped_error.WrapError(err, http.StatusBadRequest))
		return
	}

	if err := c.service.ProcessRegister(dto); err != nil {
		response.ErrorWrapped(w, err)
		return
	}

	response.OkWithMessage(w, "success")
}

func (c controller) Login(w http.ResponseWriter, r *http.Request) {
	var dto LoginDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.ErrorWrapped(w, wraped_error.WrapError(err, http.StatusBadRequest))
		return
	}

	loginResponseDto, err := c.service.ProcessLogin(dto)
	if err != nil {
		response.ErrorWrapped(w, err)
		return
	}

	response.Ok(w, loginResponseDto)
}

func (c controller) Update(w http.ResponseWriter, r *http.Request) {
	// can only update name, email, password
	var dto AccountDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.ErrorWrapped(w, wraped_error.WrapError(err, http.StatusBadRequest))
		return
	}

	if err := c.service.ProcessUpdate(dto); err != nil {
		response.ErrorWrapped(w, err)
		return
	}

	response.OkWithMessage(w, "success")
}

func (c controller) GetAccount(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	accountListDto, err := c.service.ProcessGetAccount(name)
	if err != nil {
		response.ErrorWrapped(w, err)
		return
	}

	response.Ok(w, accountListDto)
}
