package files

import (
	"lending-service/config/database"
)

type RepositoryImpl struct {
}

func NewRepository() RepositoryImpl {
	return RepositoryImpl{}
}

func (r RepositoryImpl) Insert(entity File) (int, error) {
	_, err := database.Postgres.Model(&entity).Insert()
	return entity.Id, err
}

func (r RepositoryImpl) Get(queryParam FileQueryParam) (file File, err error) {
	query := database.Postgres.Model(&file).Limit(1)
	if queryParam.Id != 0 {
		err = query.Where("id = ?", queryParam.Id).Select()
		return file, err
	}
	if queryParam.AccountId != 0 {
		query.Where("account_id = ?", queryParam.AccountId)
	}
	if queryParam.Type != "" {
		query.Where("type = ?", queryParam.Type)
	}
	err = query.Select()
	return file, err
}
