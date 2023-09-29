package loan

import (
	"github.com/go-pg/pg/v10"

	"lending-service/account"
)

type RepositoryPostgresImpl struct {
	postgres *pg.DB
}

func NewPostgresRepository(postgres *pg.DB) RepositoryPostgresImpl {
	return RepositoryPostgresImpl{
		postgres: postgres,
	}
}

func (r RepositoryPostgresImpl) InsertToDb(loan Loan) error {
	_, err := r.postgres.Model(&loan).Insert()
	return err
}

func (r RepositoryPostgresImpl) GetLoansWithParameter(account account.Account, queryParameter LoanQueryParameter) ([]Loan, error) {
	var loans []Loan
	offset := queryParameter.Page * queryParameter.Limit
	query := r.postgres.Model(&loans).Where("lender = ?", account.Id).WhereOr("borrower = ?", account.Id).Order("created_at desc").Offset(offset).Limit(queryParameter.Limit)
	if queryParameter.Status != "" {
		query.Where("status = ?", queryParameter.Status)
	}
	err := query.Select()
	return loans, err
}

func (r RepositoryPostgresImpl) GetLoanById(id int) (Loan, error) {
	var loan Loan
	err := r.postgres.Model(&loan).Where("id = ?", id).Select()
	return loan, err
}

func (r RepositoryPostgresImpl) UpdateLoan(loan Loan) error {
	_, err := r.postgres.Model(&loan).WherePK().Update()
	return err
}
