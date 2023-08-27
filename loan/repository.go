package loan

type repository interface {
	InsertToDb(loan Loan) error
}
