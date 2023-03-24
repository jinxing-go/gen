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
			&cli.BoolFlag{
				Name:    "repo",
				Aliases: []string{"r"},
				Usage:   "generate repository file, eg: --repo=true",
				Value:   false,
			},
		},
		Action: func(c *cli.Context) error {
			if err := gen.Generate(c.String("table"), c.Bool("repo")); err != nil {
				color.Red.Printf("generate model error: %v", err)
			}

			return nil
		},
	}
}
