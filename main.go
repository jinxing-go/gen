package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinxing-go/gen/config"
	"github.com/jinxing-go/gen/pkg/util"
	"github.com/urfave/cli/v2"
)

func main() {
	var conf *config.Config
	if root, err := util.FindProjectRootPath(); err != nil {
		conf = config.Load(filepath.Join(root, ".gen.yml"))
	} else {
		conf = config.Load("./.gen.yml")
	}

	commands, cleanup := bootstrap(conf)
	defer cleanup()

	app := &cli.App{
		Name:     "gen",
		Usage:    "generate code",
		Version:  "v1.0.0",
		Compiled: time.Now(),
		Commands: commands,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "print debug log",
			},
		},
		Before: func(c *cli.Context) error {
			if c.Bool("debug") {
				_ = os.Setenv("DEBUG", "true")
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
