package account

type repository interface {
	insertToDb(entity Account) error
}
