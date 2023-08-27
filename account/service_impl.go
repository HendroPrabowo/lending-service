package account

import (
	"fmt"
	"net/http"
	"net/mail"
	"regexp"
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
	if err := validateEmptyString(dto.Username, constant.USERNAME); err != nil {
		return err
	}
	isWhitespacePresent := regexp.MustCompile(`\s`).MatchString(dto.Username)
	if isWhitespacePresent {
		return fmt.Errorf(constant.ERROR_MESSAGE_USERNAME_CANNOT_CONTAIN_WHITESPACE)
	}
	if err := validateEmptyString(dto.Password, constant.PASSWORD); err != nil {
		return err
	}
	if err := validateEmptyString(dto.Name, constant.NAME); err != nil {
		return err
	}
	if err := validateEmptyString(dto.Email, constant.EMAIL); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(dto.Email); err != nil {
		return fmt.Errorf("invalid email")
	}
	return nil
}

func (svc serviceImpl) validateLoginAccountDto(dto LoginDto) error {
	if err := validateEmptyString(dto.Username, constant.USERNAME); err != nil {
		return err
	}
	if err := validateEmptyString(dto.Password, constant.PASSWORD); err != nil {
		return err
	}
	isWhitespacePresent := regexp.MustCompile(`\s`).MatchString(dto.Username)
	if isWhitespacePresent {
		return fmt.Errorf(constant.ERROR_MESSAGE_USERNAME_CANNOT_CONTAIN_WHITESPACE)
	}
	return nil
}

func validateEmptyString(text, field string) (err error) {
	if text == "" {
		return fmt.Errorf(constant.ERROR_MESSAGE_CANNOT_EMPTY, field)
	}
	return
}

func (svc serviceImpl) ProcessLogin(dto LoginDto) (LoginResponseDto, *wraped_error.Error) {
	var loginResponseDto LoginResponseDto
	if err := svc.validateLoginAccountDto(dto); err != nil {
		return loginResponseDto, wraped_error.WrapError(err, http.StatusBadRequest)
	}

	account, err := svc.repository.GetByUsername(dto.Username)
	if err != nil {
		if err == pg.ErrNoRows {
			return loginResponseDto, wraped_error.WrapError(fmt.Errorf("usename not found"), http.StatusBadRequest)
		}
		return loginResponseDto, wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	password, err := password.Decrypt(account.Password)
	if err != nil {
		return loginResponseDto, wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	if dto.Password != password {
		return loginResponseDto, wraped_error.WrapError(fmt.Errorf("username atau password salah"), http.StatusBadRequest)
	}

	loginResponseDto, errWrap := svc.generateTokenJwt(account)
	if errWrap != nil {
		return loginResponseDto, wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	loginResponseDto.Name = account.Name
	loginResponseDto.Id = account.Id
	return loginResponseDto, nil
}

func (svc serviceImpl) generateTokenJwt(account Account) (LoginResponseDto, *wraped_error.Error) {
	loginResponseDto := LoginResponseDto{}
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
		return loginResponseDto, wraped_error.WrapError(err, http.StatusInternalServerError)
	}
	loginResponseDto.Token = signedToken
	return loginResponseDto, nil
}

func (svc serviceImpl) ProcessUpdate(dto AccountDto) *wraped_error.Error {
	entity, err := svc.repository.GetByUsername(dto.Username)
	if err != nil {
		return wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	copier.CopyWithOption(&entity, &dto, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	entity.Password = password.Encrypt(entity.Password)
	entity.UpdatedAt = time_now.Wib().Format(constant.TIMEFORMAT)

	if err := svc.repository.Update(entity); err != nil {
		return wraped_error.WrapError(err, http.StatusInternalServerError)
	}
	return nil
}

func (svc serviceImpl) ProcessGetAccount(name string) ([]AccountListDto, *wraped_error.Error) {
	accountListEntity, err := svc.repository.GetByName(name)
	if err != nil {
		return nil, wraped_error.WrapError(err, http.StatusInternalServerError)
	}
	accountListDto := []AccountListDto{}
	copier.Copy(&accountListDto, &accountListEntity)
	return accountListDto, nil
}
