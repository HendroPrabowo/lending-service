//+build wireinject

package account

import (
	"github.com/google/wire"
)

func InitializeAccount() (routes, error) {
	wire.Build(
		newRoutes,
		newController,
		newService,
		NewRepository,
		NewMiddleware,

		wire.Bind(new(Service), new(serviceImpl)),
		wire.Bind(new(Repository), new(RepositoryImpl)),
	)
	return routes{}, nil
}
