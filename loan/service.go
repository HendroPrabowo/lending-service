package loan

import (
	"lending-service/account"
	"lending-service/utility/wraped_error"
)

type Service interface {
	ProceddAddLoan(loanDto LoanDto, account account.Account) *wraped_error.Error
	ProceddGetLoan(account account.Account) ([]LoanDto, *wraped_error.Error)
}