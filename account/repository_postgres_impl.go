package account

import (
	"github.com/go-pg/pg/v10"
)

type RepositoryPostgresImpl struct {
	postgres *pg.DB
}

func NewPostgresRepository(postgres *pg.DB) RepositoryPostgresImpl {
	return RepositoryPostgresImpl{
		postgres: postgres,
	}
}

func (r RepositoryPostgresImpl) InsertToDb(entity Account) error {
	_, err := r.postgres.Model(&entity).Insert()
	return err
}

func (r RepositoryPostgresImpl) GetByUsername(username string) (Account, error) {
	account := new(Account)
	err := r.postgres.Model(account).Where("username = ?", username).Select()
	return *account, err
}

func (r RepositoryPostgresImpl) Update(account Account) error {
	_, err := r.postgres.Model(&account).
		Column("password", "name", "email", "updated_at").
		WherePK().
		Update()
	return err
}

func (r RepositoryPostgresImpl) GetByName(name string) ([]Account, error) {
	var accounts []Account
	err := r.postgres.Model(&accounts).Where("name LIKE ?", name+"%").Order("name ASC").Select()
	return accounts, err
}

func (r RepositoryPostgresImpl) GetById(id int) (Account, error) {
	var account Account
	err := r.postgres.Model(&account).Where("id = ?", id).Select()
	return account, err
}
