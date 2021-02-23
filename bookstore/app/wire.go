// +build wireinject

package app

import (
	"bookstore/bookstore/config"
	"bookstore/bookstore/controller"
	"bookstore/bookstore/dao"
	"bookstore/bookstore/mysql"
	"bookstore/bookstore/service"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	dao.ProviderSet,
	service.ProviderSet,
	controller.ProviderSet,
	config.NewConfig,
	mysql.NewDB,
)

func CreateApp() (*Application, error) {
	panic(wire.Build(NewApplication, ProviderSet))
}
