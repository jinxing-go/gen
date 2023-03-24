package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gookit/color"
	"github.com/jinxing-go/gen/config"
	"github.com/jinxing-go/gen/pkg/util"
	"github.com/urfave/cli/v2"
)

func main() {

	root, err := util.FindProjectRootPath(".gen.yml")
	if err != nil {
		color.Errorf("cannot find project root path: %s", err)
	}

	conf := config.Load(filepath.Join(root, ".gen.yml"))
	if conf.ProjectPath == "" {
		conf.ProjectPath = root
	}

	commands, cleanup := bootstrap(conf)
	defer cleanup()

	app := &cli.App{
		Name:     "gen",
		Usage:    "generate code",
		Version:  "v1.0.0",
		Compiled: time.Now(),
		Commands: commands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
