package loan

import (
	"fmt"
	"net/http"
	"strconv"

	"lending-service/account"
	"lending-service/constant"
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

func (svc serviceImpl) ProceddAddLoan(loanDto LoanDto, account account.Account) *wraped_error.Error {
	if err := svc.validateLoanDto(loanDto); err != nil {
		return wraped_error.WrapError(err, http.StatusBadRequest)
	}

	loan := svc.constructLoan(loanDto, account)
	if err := svc.repository.InsertToDb(loan); err != nil {
		return wraped_error.WrapError(err, http.StatusInternalServerError)
	}
	return nil
}

func (svc serviceImpl) constructLoan(dto LoanDto, account account.Account) Loan {
	amount, _ := strconv.Atoi(dto.Amount)
	loan := Loan{
		Lender:      account.Id,
		Borrower:    dto.Borrower,
		Amount:      amount,
		Status:      constant.LOAN_STATUS_UNPAID,
		Description: dto.Description,
	}
	return loan
}

func (svc serviceImpl) validateLoanDto(dto LoanDto) error {
	if dto.Borrower <= 0 {
		return fmt.Errorf("borrower cannot empty")
	}
	if dto.Description == "" {
		return fmt.Errorf("description cannot empty")
	}
	if dto.Amount == "" {
		return fmt.Errorf("ammount cannot empty")
	}
	amount, err := strconv.Atoi(dto.Amount)
	if err != nil {
		return fmt.Errorf("amount must in number")
	}
	if amount <= 0 {
		return fmt.Errorf("ammount cannot empty")
	}
	return nil
}
