package loan

type LoanDto struct {
	Lender       int    `json:"lender"`
	LenderName   string `json:"lender_name"`
	Borrower     int    `json:"borrower"`
	BorrowerName string `json:"borrower_name"`
	Description  string `json:"description"`
	Amount       int    `json:"amount"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
