package account

import (
	"fmt"
	"net/http"
	"net/mail"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/copier"

	"lending-service/constant"
	"lending-service/utility/password"
	"lending-service/utility/time_now"
	"lending-service/utility/wraped_error"
)

type serviceImpl struct {
	repository Repository
}

func newService(repository Repository) serviceImpl {
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

	if err := svc.repository.InsertToDb(entity); err != nil {
		return wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	return nil
}

func (svc serviceImpl) validateAccountDto(dto AccountDto) error {
	if err := svc.validateLoginAccountDto(dto); err != nil {
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

func (svc serviceImpl) validateLoginAccountDto(dto AccountDto) error {
	if err := validateEmptyString(dto.Username, "username"); err != nil {
		return err
	}
	if err := validateEmptyString(dto.Password, "password"); err != nil {
		return err
	}
	return nil
}

func validateEmptyString(text, field string) (err error) {
	if text == "" {
		return fmt.Errorf(constant.ErrorMessageCannotEmpty, field)
	}
	return
}

func (svc serviceImpl) ProcessLogin(dto AccountDto) (LoginDto, *wraped_error.Error) {
	var loginDto LoginDto
	if err := svc.validateLoginAccountDto(dto); err != nil {
		return loginDto, wraped_error.WrapError(err, http.StatusBadRequest)
	}

	account, err := svc.repository.GetByUsername(dto.Username)
	if err != nil {
		if err == pg.ErrNoRows {
			return loginDto, wraped_error.WrapError(fmt.Errorf("usename not found"), http.StatusBadRequest)
		}
		return loginDto, wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	loginDto, errWrap := svc.generateTokenJwt(account)
	if errWrap != nil {
		return loginDto, wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	return loginDto, nil
}

func (svc serviceImpl) generateTokenJwt(account Account) (LoginDto, *wraped_error.Error) {
	loginDto := LoginDto{}
	claim := Claims{
		Account: account,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			Issuer:    account.Name,
		},
	}

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := tokenClaim.SignedString(constant.JWT_KEY)
	if err != nil {
		return loginDto, wraped_error.WrapError(err, http.StatusInternalServerError)
	}
	loginDto.Token = signedToken
	return loginDto, nil
}

func (svc serviceImpl) ProcessUpdate(dto AccountDto) *wraped_error.Error {
	entity, err := svc.repository.GetByUsername(dto.Username)
	if err != nil {
		return wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	copier.CopyWithOption(&entity, &dto, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	entity.UpdatedAt = time_now.Wib().Format(constant.TIMEFORMAT)

	if err := svc.repository.Update(entity); err != nil {
		return wraped_error.WrapError(err, http.StatusInternalServerError)
	}
	return nil
}
