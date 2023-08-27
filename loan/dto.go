package loan

type LoanDto struct {
	Borrower    int    `json:"borrower"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
}
