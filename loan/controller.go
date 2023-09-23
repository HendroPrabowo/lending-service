package loan

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"lending-service/account"
	"lending-service/constant"
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

func (c controller) GetLoan(w http.ResponseWriter, r *http.Request) {
	queryParameter, err := c.buildLoanQueryParameter(r)
	if err != nil {
		response.ErrorWrapped(w, err)
		return
	}

	account := r.Context().Value(constant.ACCOUNT).(account.Account)
	loansDto, err := c.service.ProceedGetLoans(account, queryParameter)
	if err != nil {
		response.ErrorWrapped(w, err)
		return
	}
	response.Ok(w, loansDto)
}

func (c controller) buildLoanQueryParameter(r *http.Request) (LoanQueryParameter, *wraped_error.Error) {
	var loanQueryParameter LoanQueryParameter

	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	status := r.URL.Query().Get("status")

	if page == "" {
		page = "0"
	}

	if limit == "" {
		limit = "25"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return loanQueryParameter, wraped_error.WrapError(fmt.Errorf("invalid value page"), http.StatusBadRequest)
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return loanQueryParameter, wraped_error.WrapError(fmt.Errorf("invalid value limit"), http.StatusBadRequest)
	}

	loanQueryParameter.Page = pageInt
	loanQueryParameter.Limit = limitInt
	loanQueryParameter.Status = status
	return loanQueryParameter, nil
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

func (c controller) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	var loanStatusDto LoanStatusDto
	if err := json.NewDecoder(r.Body).Decode(&loanStatusDto); err != nil {
		response.ErrorWithMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	account := r.Context().Value(constant.ACCOUNT).(account.Account)
	if err := c.service.ChangeLoanStatus(account, loanStatusDto); err != nil {
		response.ErrorWrapped(w, err)
		return
	}

	response.OkWithMessage(w, "success")
}
