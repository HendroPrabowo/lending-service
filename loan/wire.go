//+build wireinject

package loan

import (
	"github.com/google/wire"

	"lending-service/account"
)

func InitializeLoan() (routes, error) {
	wire.Build(
		newRoutes,
		account.NewMiddleware,
		newController,
		newService,
		newRepository,
		account.NewRepository,

		wire.Bind(new(Service), new(serviceImpl)),
		wire.Bind(new(repository), new(repositoryImpl)),
		wire.Bind(new(account.Repository), new(account.RepositoryImpl)),
	)
	return routes{}, nil
}
