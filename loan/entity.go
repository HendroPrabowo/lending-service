package loan

type Loan struct {
	tableName   struct{} `pg:"loan"`
	Id          int
	Lender      int
	Borrower    int
	Amount      int
	Status      string
	Description string
	CreatedAt   string
	UpdatedAt   string
}
