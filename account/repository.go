package account

type Repository interface {
	InsertToDb(entity Account) error
	GetByUsername(username string) (Account, error)
	Update(account Account) error
}
