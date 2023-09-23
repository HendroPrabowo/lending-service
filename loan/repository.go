package loan

import "lending-service/account"

type repository interface {
	InsertToDb(loan Loan) error
	GetLoans(account account.Account) ([]Loan, error)
	GetLoanById(id int) (Loan, error)
	UpdateLoan(loan Loan) error
}
