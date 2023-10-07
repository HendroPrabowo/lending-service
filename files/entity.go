package files

type File struct {
	tableName   struct{} `pg:"files"`
	Id          int
	AccountId   int
	Type        string
	ContentType string
	File        string
	CreatedAt   string
	UpdatedAt   string
}
