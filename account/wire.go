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

		wire.Bind(new(service), new(serviceImpl)),
		wire.Bind(new(repository), new(repositoryImpl)),
	)
	return routes{}, nil
}
