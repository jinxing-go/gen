package main

import (
	"fmt"
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

	root, err := util.FindProjectRootPath(".gen.yml")
	if err != nil {
		fmt.Println("cannot find project root path", err)
	}

	fmt.Println("root:", root)
	conf := config.Load(filepath.Join(root, ".gen.yml"))
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

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
