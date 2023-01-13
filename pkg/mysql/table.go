package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/jinxing-go/gen/pkg/util"
)

const (
	KeyIndex      = "MUL" // 普通索引
	KeyPrimaryKey = "PRI" // 主键ID
	KeyUnique     = "UNI" // 唯一索引

	TypeDate      = "date"
	TypeTime      = "time"
	TypeYear      = "year"
	TypeDatetime  = "datetime"
	TypeTimestamp = "timestamp"
)

type TableResult struct {
	Field      string         `json:"field"`     // 字段名称
	Type       string         `json:"type"`      // 字段类型
	Collation  sql.NullString `json:"collation"` // 字符
	Null       sql.NullString `json:"null"`      // 是否为空
	Key        sql.NullString `json:"key"`       // 主键
	Default    sql.NullString `json:"default"`   // 默认值
	Extra      sql.NullString `json:"extra"`
	Privileges sql.NullString `json:"privileges"`
	Comment    sql.NullString `json:"comment"` // 字段注释
}

type TableColumn struct {
	StudlyName  string `json:"studly_name"`            // 大驼峰命名
	Field       string `json:"field"`                  // 字段名称
	Type        string `json:"type,omitempty"`         // 字段类型
	Null        string `json:"null,omitempty"`         // 是否为空
	Key         string `json:"key,omitempty"`          // 主键
	Comment     string `json:"comment,omitempty"`      // 字段注释
	MappingType string `json:"mapping_type,omitempty"` // 映射类型
}

type Table struct {
	StudlyName   string         `json:"studly_name"`   // 大驼峰命名
	DBName       string         `json:"db_name"`       // 库名称
	TableName    string         `json:"table_name"`    // 表名
	Comment      string         `json:"comment"`       // 表注释
	PrimaryKey   *TableColumn   `json:"primary_key"`   // 主键
	Columns      []*TableColumn `json:"columns"`       // 字段信息
	PrimaryField string         `json:"primary_field"` // 主键ID
	HasTimestamp bool           `json:"has_timestamp"` // 是否有时间戳
}

var defaultTypeMap = map[string]string{
	"int":        "int",
	"tinyint":    "int",
	"smallint":   "int",
	"mediumint":  "int",
	"enum":       "int",
	"bigint":     "int64",
	"char":       "string",
	"varchar":    "string",
	"json":       "string",
	"year":       "time.Time",
	"timestamp":  "time.Time",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"text":       "string",
	"mediumtext": "string",
	"longtext":   "string",
	"double":     "float64",
	"decimal":    "float64",
	"float":      "float64",
}

func NewTable(db *sql.DB, dbName, tableName string, mapType map[string]string) (*Table, error) {

	table := &Table{
		StudlyName: util.Studly(tableName),
		DBName:     dbName,
		TableName:  tableName,
		PrimaryKey: &TableColumn{},
	}

	if err := table.queryComment(db); err != nil {
		return nil, err
	}

	if err := table.queryColumns(db); err != nil {
		return nil, err
	}

	for _, column := range table.Columns {
		if util.Includes(column.Type, []string{TypeDate, TypeTime, TypeYear, TypeDatetime, TypeTimestamp}) {
			table.HasTimestamp = true
		}

		column.StudlyName = util.Studly(column.Field)
		column.MappingType = getType(column.Type, mapType)
		if column.Key == KeyPrimaryKey {
			table.PrimaryKey = column
			table.PrimaryField = column.Field
		}
	}

	return table, nil
}

func (t *Table) queryComment(db *sql.DB) error {
	rows, err := db.Query(fmt.Sprintf("SELECT TABLE_COMMENT FROM information_schema.TABLES WHERE `TABLE_SCHEMA` = '%s' AND TABLE_NAME = '%s'", t.DBName, t.TableName))
	if err != nil {
		return fmt.Errorf("查询表注释失败: %s", err)
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&t.Comment); err != nil {
			return fmt.Errorf("解析注释信息失败: %s", err)
		}
	}

	return nil
}

func (t *Table) queryColumns(db *sql.DB) error {
	rows, err := db.Query(fmt.Sprintf("SHOW FULL COLUMNS FROM `%s`", t.TableName))
	if err != nil {
		return fmt.Errorf("查询SQL错误: %s", err)
	}

	defer rows.Close()

	t.Columns = make([]*TableColumn, 0)
	for rows.Next() {
		tmpValue := &TableResult{}
		if err := rows.Scan(
			&tmpValue.Field,
			&tmpValue.Type,
			&tmpValue.Collation,
			&tmpValue.Null,
			&tmpValue.Key,
			&tmpValue.Default,
			&tmpValue.Extra,
			&tmpValue.Privileges,
			&tmpValue.Comment,
		); err != nil {
			log.Printf("扫描%s错误: %s\n", tmpValue.Field, err)
		}

		t.Columns = append(t.Columns, &TableColumn{
			Field:   tmpValue.Field,
			Type:    toType(tmpValue.Type),
			Null:    tmpValue.Null.String,
			Key:     tmpValue.Key.String,
			Comment: tmpValue.Comment.String,
		})
	}

	return nil
}

func toType(s string) string {
	if strings.Contains(s, "(") {
		return strings.Split(s, "(")[0]
	}

	if strings.Contains(s, " ") {
		return strings.Split(s, " ")[0]
	}

	return s
}

func getType(typeName string, mapType map[string]string) string {
	if v, ok := mapType[typeName]; ok {
		return v
	}

	if v, ok := defaultTypeMap[typeName]; ok {
		return v
	}

	return "string"
}
