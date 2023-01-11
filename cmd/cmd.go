package cmd

import (
	"reflect"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

type Commands struct {
	ModelCommand ModelCommand
}

func NewCommands(cmd *Commands) cli.Commands {
	var commands cli.Commands
	v := reflect.Indirect(reflect.ValueOf(cmd))
	ct := reflect.TypeOf(&cli.Command{})
	for i := 0; i < v.NumField(); i++ {
		commands = append(commands, v.Field(i).Convert(ct).Interface().(*cli.Command))
	}

	return commands
}

var ProviderSet = wire.NewSet(
	NewModelCommand,
	NewCommands,
	wire.Struct(new(Commands), "*"),
)
