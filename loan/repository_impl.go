package loan

import (
	"lending-service/account"
	"lending-service/config/database"
)

type repositoryImpl struct {
}

func newRepository() repositoryImpl {
	return repositoryImpl{}
}

func (r repositoryImpl) InsertToDb(loan Loan) error {
	_, err := database.Postgres.Model(&loan).Insert()
	return err
}

func (r repositoryImpl) GetLoan(account account.Account) ([]Loan, error) {
	var loans []Loan
	err := database.Postgres.Model(&loans).Where("lender = ?", account.Id).WhereOr("borrower = ?", account.Id).Select()
	return loans, err
}