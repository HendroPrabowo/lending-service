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
	repository        repository
	accountRepository account.Repository
}

func newService(repository repository, accountRepository account.Repository) serviceImpl {
	return serviceImpl{
		repository:        repository,
		accountRepository: accountRepository,
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

func (svc serviceImpl) ProceedGetLoans(account account.Account, queryParameter LoanQueryParameter) ([]LoanDto, *wraped_error.Error) {
	// initial memory cache name
	nameMap := map[int]string{account.Id: account.Name}
	loanDtos := []LoanDto{}

	loansEntity, err := svc.repository.GetLoansWithParameter(account, queryParameter)
	if err != nil {
		return nil, wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	for _, loan := range loansEntity {
		lenderName, err := svc.getAndSetNameToMap(loan.Lender, nameMap)
		if err != nil {
			return nil, wraped_error.WrapError(err, http.StatusInternalServerError)
		}
		borrowerName, err := svc.getAndSetNameToMap(loan.Borrower, nameMap)
		if err != nil {
			return nil, wraped_error.WrapError(err, http.StatusInternalServerError)
		}
		var loanDto LoanDto
		copier.Copy(&loanDto, &loan)
		loanDto.BorrowerName = borrowerName
		loanDto.LenderName = lenderName
		loanDtos = append(loanDtos, loanDto)
	}
	return loanDtos, nil
}

func (svc serviceImpl) getAndSetNameToMap(accountId int, nameMap map[int]string) (string, error) {
	name, ok := nameMap[accountId]
	if !ok {
		account, err := svc.accountRepository.GetById(accountId)
		if err != nil {
			return "", err
		}
		nameMap[accountId] = account.Name
		name = account.Name
	}
	return name, nil
}

func (svc serviceImpl) ChangeLoanStatus(account account.Account, dto LoanStatusDto) *wraped_error.Error {
	loan, err := svc.repository.GetLoanById(dto.Id)
	if err != nil {
		return wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	valid, ok := loanStatusMap[dto.Status]
	if !ok || !valid {
		return wraped_error.WrapError(fmt.Errorf("invalid status : %s", dto.Status), http.StatusBadRequest)
	}

	if loan.Lender != account.Id && loan.Borrower != account.Id {
		return wraped_error.WrapError(fmt.Errorf("this loan is not belong to this user"), http.StatusBadRequest)
	}

	loan.Status = dto.Status
	if err = svc.repository.UpdateLoan(loan); err != nil {
		return wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	return nil
}
