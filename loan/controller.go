package loan

import (
	"encoding/json"
	"net/http"

	"lending-service/account"
	"lending-service/constant"
	"lending-service/utility/response"
)

type controller struct {
	service Service
}

func newController(service Service) controller {
	return controller{
		service: service,
	}
}

func (c controller) AddLoan(w http.ResponseWriter, r *http.Request) {
	var loanDto LoanDto
	if err := json.NewDecoder(r.Body).Decode(&loanDto); err != nil {
		response.ErrorWithMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	account := r.Context().Value(constant.ACCOUNT).(account.Account)
	if err := c.service.ProceddAddLoan(loanDto, account); err != nil {
		response.ErrorWrapped(w, err)
		return
	}

	response.OkWithMessage(w, "success")
}
