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
		newRepository,
		newMiddleware,

		wire.Bind(new(Service), new(serviceImpl)),
		wire.Bind(new(Repository), new(repositoryImpl)),
	)
	return routes{}, nil
}
