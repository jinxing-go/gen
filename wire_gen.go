// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jinxing-go/gen/cmd"
	"github.com/jinxing-go/gen/config"
	"github.com/jinxing-go/gen/model"
	"github.com/urfave/cli/v2"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func bootstrap(conf *config.Config) (cli.Commands, func()) {
	db := config.NewDefaultDB(conf)
	generator := &model.Generator{
		DB:     db,
		Config: conf,
	}
	modelCommand := cmd.NewModelCommand(generator)
	commands := &cmd.Commands{
		ModelCommand: modelCommand,
	}
	cliCommands := cmd.NewCommands(commands)
	return cliCommands, func() {
	}
}

// wire.go:

var providerSet = wire.NewSet(config.NewDefaultDB, wire.Struct(new(model.Generator), "*"), cmd.ProviderSet)
