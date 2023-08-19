package account

import (
	"lending-service/utility/wraped_error"
)

type Service interface {
	ProcessRegister(dto AccountDto) *wraped_error.Error
	ProcessLogin(dto LoginDto) (LoginResponseDto, *wraped_error.Error)
	ProcessUpdate(dto AccountDto) *wraped_error.Error
}
