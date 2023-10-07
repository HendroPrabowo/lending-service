package files

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

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

func (c controller) Download(w http.ResponseWriter, r *http.Request) {
	queryParam, errWrap := c.buildQueryParam(r)
	if errWrap != nil {
		response.ErrorWrapped(w, errWrap)
		return
	}

	byteFile, contentType, errWrap := c.service.proceedGetFile(queryParam)
	if errWrap != nil {
		response.ErrorWrapped(w, errWrap)
		return
	}

	w.Header().Set("Content-Type", contentType)
	_, err := io.Copy(w, strings.NewReader(string(byteFile)))
	if err != nil {
		response.ErrorWrapped(w, wraped_error.WrapError(err, http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c controller) buildQueryParam(r *http.Request) (FileQueryParam, *wraped_error.Error) {
	queryParam := FileQueryParam{}

	idStr := r.URL.Query().Get("id")
	accountIdStr := r.URL.Query().Get("account_id")
	fileType := r.URL.Query().Get("type")

	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return queryParam, wraped_error.WrapError(fmt.Errorf("id should in integer"), http.StatusBadRequest)
		}
		queryParam.Id = id
	}

	if accountIdStr != "" {
		accountId, err := strconv.Atoi(accountIdStr)
		if err != nil {
			return queryParam, wraped_error.WrapError(fmt.Errorf("account_id should in integer"), http.StatusBadRequest)
		}
		queryParam.AccountId = accountId
	}
	queryParam.Type = fileType
	return queryParam, nil
}

func (c controller) Upload(w http.ResponseWriter, r *http.Request) {
	account := r.Context().Value(constant.ACCOUNT).(account.Account)
	file, fileHeader, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		response.ErrorWrapped(w, wraped_error.WrapError(err, http.StatusBadRequest))
		return
	}

	resposne, errWrap := c.service.proceedSaveFile(file, fileHeader, account)
	if errWrap != nil {
		response.ErrorWrapped(w, errWrap)
		return
	}

	response.Ok(w, resposne)
}
