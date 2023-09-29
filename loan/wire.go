//+build wireinject

package loan

import (
	"github.com/google/wire"

	"lending-service/account"
	"lending-service/config/database"
)

func InitializeLoanWithPostgres() (routes, error) {
	wire.Build(
		newRoutes,
		account.NewMiddleware,
		newController,
		newService,
		NewPostgresRepository,
		account.NewPostgresRepository,
		database.InitPostgreOrm,

		wire.Bind(new(Service), new(serviceImpl)),
		wire.Bind(new(Repository), new(RepositoryPostgresImpl)),
		wire.Bind(new(account.Repository), new(account.RepositoryPostgresImpl)),
	)
	return routes{}, nil
}

func InitializeLoanWithMysql() (routes, error) {
	wire.Build(
		newRoutes,
		account.NewMiddleware,
		newController,
		newService,
		NewMysqlRepository,
		account.NewMysqlRepository,
		database.InitMysql,

		wire.Bind(new(Service), new(serviceImpl)),
		wire.Bind(new(Repository), new(RepositoryMysqlImpl)),
		wire.Bind(new(account.Repository), new(account.RepositoryMysqlImpl)),
	)
	return routes{}, nil
}
