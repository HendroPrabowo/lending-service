package loan

import (
	"gorm.io/gorm"

	"lending-service/account"
)

type RepositoryMysqlImpl struct {
	mysql *gorm.DB
}

func NewMysqlRepository(mysql *gorm.DB) RepositoryMysqlImpl {
	return RepositoryMysqlImpl{
		mysql: mysql,
	}
}

func (r RepositoryMysqlImpl) InsertToDb(loan Loan) error {
	return nil
}

func (r RepositoryMysqlImpl) GetLoansWithParameter(account account.Account, queryParameter LoanQueryParameter) ([]Loan, error) {
	return nil, nil
}

func (r RepositoryMysqlImpl) GetLoanById(id int) (Loan, error) {
	return Loan{}, nil
}

func (r RepositoryMysqlImpl) UpdateLoan(loan Loan) error {
	return nil
}
