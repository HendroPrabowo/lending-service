//+build wireinject

package account

import (
	"github.com/google/wire"

	"lending-service/config/database"
)

func InitializeAccountWithPostgres() (routes, error) {
	wire.Build(
		newRoutes,
		newController,
		newService,
		NewPostgresRepository,
		NewMiddleware,
		database.InitPostgreOrm,

		wire.Bind(new(Service), new(serviceImpl)),
		wire.Bind(new(Repository), new(RepositoryPostgresImpl)),
	)
	return routes{}, nil
}

func InitializeAccountWithMysql() (routes, error) {
	wire.Build(
		newRoutes,
		newController,
		newService,
		NewMysqlRepository,
		NewMiddleware,
		database.InitMysql,

		wire.Bind(new(Service), new(serviceImpl)),
		wire.Bind(new(Repository), new(RepositoryMysqlImpl)),
	)
	return routes{}, nil
}
