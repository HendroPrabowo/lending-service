package loan

type LoanDto struct {
	Lender      int    `json:"lender"`
	Borrower    int    `json:"borrower"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
