package account

import (
	"lending-service/config/database"
)

type RepositoryImpl struct {
}

func NewRepository() RepositoryImpl {
	return RepositoryImpl{}
}

func (r RepositoryImpl) InsertToDb(entity Account) error {
	_, err := database.Postgres.Model(&entity).Insert()
	return err
}

func (r RepositoryImpl) GetByUsername(username string) (Account, error) {
	account := new(Account)
	err := database.Postgres.Model(account).Where("username = ?", username).Select()
	return *account, err
}

func (r RepositoryImpl) Update(account Account) error {
	_, err := database.Postgres.Model(&account).
		Column("password", "name", "email", "updated_at").
		WherePK().
		Update()
	return err
}

func (r RepositoryImpl) GetByName(name string) ([]Account, error) {
	var accounts []Account
	err := database.Postgres.Model(&accounts).Where("name LIKE ?", name+"%").Order("name ASC").Select()
	return accounts, err
}

func (r RepositoryImpl) GetById(id int) (Account, error) {
	var account Account
	err := database.Postgres.Model(&account).Where("id = ?", id).Select()
	return account, err
}
