package loan

import "lending-service/config/database"

type repositoryImpl struct {
}

func newRepository() repositoryImpl {
	return repositoryImpl{}
}

func (r repositoryImpl) InsertToDb(loan Loan) error {
	_, err := database.Postgres.Model(&loan).Insert()
	return err
}
