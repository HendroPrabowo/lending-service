package account

import "lending-service/config/database"

type repositoryImpl struct {
}

func newRepository() repositoryImpl {
	return repositoryImpl{}
}

func (r repositoryImpl) insertToDb(entity Account) error {
	_, err := database.Orm.Model(&entity).Insert()
	return err
}
