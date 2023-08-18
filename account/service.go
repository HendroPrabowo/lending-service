package account

import (
	"lending-service/utility/wraped_error"
)

type service interface {
	ProcessRegister(dto AccountDto) *wraped_error.Error
}