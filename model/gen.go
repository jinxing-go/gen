package model

import (
	"database/sql"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gookit/color"
	"github.com/jinxing-go/gen/config"
	"github.com/jinxing-go/gen/pkg/mysql"
	"github.com/jinxing-go/gen/pkg/util"
)

//go:embed template.gohtml
var tpl string

type Generator struct {
	DB     *sql.DB
	Config *config.Config
}

func (m *Generator) Generate(tableName string) error {
	// 处理查询表数据信息
	table, err := mysql.NewTable(m.DB, m.Config.DB.Name, tableName, m.Config.Model.Types)
	if err != nil {
		return err
	}

	tempValue := tpl
	if m.Config.Model.Template != "" {
		body, err := os.ReadFile(m.Config.Model.Template)
		if err != nil {
			return fmt.Errorf("read template file %s error: %w", m.Config.Model.Template, err)
		}

		tempValue = string(body)
	}

	// 处理为文件
	var content string
	if content, err = util.Parse(tempValue, table); err != nil {
		return err
	}

	// 写入目录
	filename := filepath.Join(m.Config.Model.Dirname, fmt.Sprintf("%s.go", table.TableName))
	if err := util.WriteFile(filename, content); err != nil {
		return err
	}

	color.Success.Printf("%s\n", content)
	return nil
}
