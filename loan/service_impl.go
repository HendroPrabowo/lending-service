package loan

import (
	"fmt"
	"net/http"

	"lending-service/account"
	"lending-service/constant"
	"lending-service/utility/wraped_error"

	"github.com/jinzhu/copier"
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
	loan := Loan{
		Lender:      account.Id,
		Borrower:    dto.Borrower,
		Amount:      dto.Amount,
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
	if dto.Amount <= 0 {
		return fmt.Errorf("ammount cannot empty")
	}
	return nil
}

func (svc serviceImpl) ProceddGetLoan(account account.Account) ([]LoanDto, *wraped_error.Error) {
	loansEntity, err := svc.repository.GetLoan(account)
	if err != nil {
		return nil, wraped_error.WrapError(err, http.StatusInternalServerError)
	}
	var loans []LoanDto
	copier.Copy(&loans, &loansEntity)
	return loans, nil
}
