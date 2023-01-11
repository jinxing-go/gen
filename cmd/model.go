package cmd

import (
	"github.com/gookit/color"
	"github.com/jinxing-go/gen/model"
	"github.com/urfave/cli/v2"
)

type ModelCommand *cli.Command

func NewModelCommand(gen *model.Generator) ModelCommand {
	return &cli.Command{
		Name:  "model",
		Usage: "generate model file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "table",
				Aliases:  []string{"t"},
				Usage:    "table name, default all tables, eg: --table=orders",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			if err := gen.Generate(c.String("table")); err != nil {
				color.Red.Printf("generate model error: %v", err)
			}

			return nil
		},
	}
}
