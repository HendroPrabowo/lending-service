package account

import (
	"gorm.io/gorm"
)

type RepositoryMysqlImpl struct {
	mysql *gorm.DB
}

func NewMysqlRepository(mysql *gorm.DB) RepositoryMysqlImpl {
	return RepositoryMysqlImpl{
		mysql: mysql,
	}
}

func (r RepositoryMysqlImpl) InsertToDb(entity Account) error {
	result := r.mysql.Create(entity)
	return result.Error
}

func (r RepositoryMysqlImpl) GetByUsername(username string) (Account, error) {
	return Account{}, nil
}

func (r RepositoryMysqlImpl) Update(account Account) error {
	return nil
}

func (r RepositoryMysqlImpl) GetByName(name string) ([]Account, error) {

	return nil, nil
}

func (r RepositoryMysqlImpl) GetById(id int) (Account, error) {
	return Account{}, nil
}
