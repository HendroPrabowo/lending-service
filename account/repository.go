package account

type Repository interface {
	InsertToDb(entity Account) error
	GetByUsername(username string) (Account, error)
	Update(account Account) error
	GetByName(name string) ([]Account, error)
	GetById(id int) (Account, error)
}
