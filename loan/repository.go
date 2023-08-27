package loan

import "lending-service/account"

type repository interface {
	InsertToDb(loan Loan) error
	GetLoan(account account.Account) ([]Loan, error)
}
