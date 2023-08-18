package account

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/jinzhu/copier"

	"lending-service/constant"
	"lending-service/utility/password"
	"lending-service/utility/wraped_error"
)

type serviceImpl struct {
	repository repository
}

func newService(repository repository) serviceImpl {
	return serviceImpl{
		repository: repository,
	}
}

func (svc serviceImpl) ProcessRegister(dto AccountDto) *wraped_error.Error {
	if err := svc.validateAccountDto(dto); err != nil {
		return wraped_error.WrapError(err, http.StatusBadRequest)
	}

	entity := Account{}
	copier.Copy(&entity, &dto)
	entity.Password = password.Encrypt(entity.Password)

	if err := svc.repository.insertToDb(entity); err != nil {
		return wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	return nil
}

func (svc serviceImpl) validateAccountDto(dto AccountDto) error {
	if err := validateEmptyString(dto.Username, "username"); err != nil {
		return err
	}
	if err := validateEmptyString(dto.Password, "password"); err != nil {
		return err
	}
	if err := validateEmptyString(dto.Name, "name"); err != nil {
		return err
	}
	if err := validateEmptyString(dto.Email, "email"); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(dto.Email); err != nil {
		return fmt.Errorf("invalid email")
	}
	return nil
}

func validateEmptyString(text, field string) (err error) {
	if text == "" {
		return fmt.Errorf(constant.ErrorMessageCannotEmpty, field)
	}
	return
}
