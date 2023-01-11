//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jinxing-go/gen/cmd"
	"github.com/jinxing-go/gen/config"
	"github.com/jinxing-go/gen/model"
	"github.com/urfave/cli/v2"
)

var providerSet = wire.NewSet(
	config.NewDefaultDB,
	wire.Struct(new(model.Generator), "*"),
	cmd.ProviderSet,
)

func bootstrap(conf *config.Config) (cli.Commands, func()) {
	panic(wire.Build(providerSet))
}
