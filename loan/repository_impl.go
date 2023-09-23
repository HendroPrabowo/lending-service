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

func (r repositoryImpl) GetLoansWithParameter(account account.Account, queryParameter LoanQueryParameter) ([]Loan, error) {
	var loans []Loan
	offset := queryParameter.Page * queryParameter.Limit
	query := database.Postgres.Model(&loans).Where("lender = ?", account.Id).WhereOr("borrower = ?", account.Id).Order("created_at desc").Offset(offset).Limit(queryParameter.Limit)
	if queryParameter.Status != "" {
		query.Where("status = ?", queryParameter.Status)
	}
	err := query.Select()
	return loans, err
}

func (r repositoryImpl) GetLoanById(id int) (Loan, error) {
	var loan Loan
	err := database.Postgres.Model(&loan).Where("id = ?", id).Select()
	return loan, err
}

func (r repositoryImpl) UpdateLoan(loan Loan) error {
	_, err := database.Postgres.Model(&loan).WherePK().Update()
	return err
}
