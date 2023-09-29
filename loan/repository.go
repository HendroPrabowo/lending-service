package loan

import "lending-service/account"

type Repository interface {
	InsertToDb(loan Loan) error
	GetLoansWithParameter(account account.Account, queryParameter LoanQueryParameter) ([]Loan, error)
	GetLoanById(id int) (Loan, error)
	UpdateLoan(loan Loan) error
}
