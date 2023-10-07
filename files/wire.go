//+build wireinject

package files

import (
	"github.com/google/wire"

	"lending-service/account"
)

func InitializeRoutes() (routes, error) {
	wire.Build(
		newController,
		newRoutes,
		account.NewMiddleware,
		NewRepository,
		newService,

		wire.Bind(new(Service), new(serviceImpl)),
		wire.Bind(new(Repository), new(RepositoryImpl)),
	)
	return routes{}, nil
}
