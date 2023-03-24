package model

import (
	"database/sql"
	_ "embed"
	"fmt"
	"path/filepath"

	"github.com/gookit/color"
	"github.com/jinxing-go/gen/config"
	"github.com/jinxing-go/gen/pkg/mysql"
	"github.com/jinxing-go/gen/pkg/util"
)

//go:embed template.gohtml
var tpl string

//go:embed repo.gohtml
var repoTpl string

type Generator struct {
	DB     *sql.DB
	Config *config.Config
}

func (m *Generator) Generate(tableName string, genRepo bool) error {
	// 处理查询表数据信息
	table, err := mysql.NewTable(m.DB, m.Config.DB.Name, tableName, m.Config.Model.Types)
	if err != nil {
		return err
	}

	var tempValue string
	if tempValue, err = util.EstimateReadFile(m.Config.Model.Template, tpl); err != nil {
		return err
	}

	// 处理为文件
	var content string
	if content, err = util.Parse(tempValue, table); err != nil {
		return err
	}

	// 写入目录
	filename := filepath.Join(m.Config.ProjectPath, m.Config.Model.Dirname, fmt.Sprintf("%s.go", table.TableName))
	if err := util.WriteFile(filename, content); err != nil {
		return err
	}

	color.Success.Printf("gen model %s\n", filename)

	if genRepo && m.Config.Model.RepoDirname != "" {
		// 判断读取模板文件
		if tempValue, err = util.EstimateReadFile(m.Config.Model.RepoTemplate, repoTpl); err != nil {
			return err
		}

		if content, err = util.Parse(tempValue, table); err != nil {
			return err
		}

		// 写入目录
		filename = filepath.Join(m.Config.ProjectPath, m.Config.Model.RepoDirname, fmt.Sprintf("%s.go", table.TableName))
		if err := util.WriteFile(filename, content); err != nil {
			return err
		}

		color.Success.Printf("gen repo %s\n", filename)
	}

	return nil
}
