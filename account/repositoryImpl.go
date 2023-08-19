package account

import (
	"lending-service/config/database"
)

type repositoryImpl struct {
}

func newRepository() repositoryImpl {
	return repositoryImpl{}
}

func (r repositoryImpl) InsertToDb(entity Account) error {
	_, err := database.Postgres.Model(&entity).Insert()
	return err
}

func (r repositoryImpl) GetByUsername(username string) (Account, error) {
	account := new(Account)
	err := database.Postgres.Model(account).Where("username = ?", username).Select()
	return *account, err
}

func (r repositoryImpl) Update(account Account) error {
	_, err := database.Postgres.Model(&account).
		Column("password", "name", "email", "updated_at").
		WherePK().
		Update()
	return err
}
