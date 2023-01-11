# Gen

简单生成器

## 安装

```bash
go install github.com/jinxing-go/gen
```

## 配置

需要在项目目录下添加 `.gen.yml` 文件

```yaml
db:
  # 数据库连接地址
  host: 127.0.0.0
  # 数据库连接端口
  port: 3306
  # 数据库名称
  name: test
  # 数据库用户名
  username: root
  # 数据库密码
  password: "123456"
model:
  # 生成model到指定目录
  dirname: "./model"
  # 生成model使用的模板
  template: ""
  # 生成model的类型映射: 数据库类型：go类型
  types:
    int: "int32"
```

## 使用

```bash
gen model --table=table_name
```

### 模板

`Table` 字段说明

| 字段名称         | 类型               | 说明                    | 示例          |
|--------------|------------------|-----------------------|-------------|
| StudlyName   | `string`         | 根据表面生成的大驼峰名称(model名称) | UserOrders  |
| DBName       | `string`         | 数据库名称                 | test        |
| TableName    | `string`         | 表名称                   | user_orders |
| Comment      | `string`         | 表注释                   | 用户订单表       |
| HasTimestamp | `bool`           | 是否有时间戳字段              | false       |
| PrimaryField | `string`         | 主键字段名称                | id          |
| PrimaryKey   | `*TableColumn`   | 主键字段信息                |             |
| Columns      | `[]*TableColumn` | 表字段信息（包含主键字段）         |             |

`TableColumn` 字段说明

| 字段名称        | 类型       | 说明           | 示例          |
|-------------|----------|--------------|-------------|
| StudlyName  | `string` | 表字段名称大驼峰名称   | UserId      |
| Field       | `string` | 表字段名称        | user_id     |
| Type        | `string` | 表字段类型(MySQL) | int         |
| Null        | `string` | 是否为NO(MySQL) | NO          |
| Key         | `string` | 索引信息         | PRI、UNI、MUL |
| Comment     | `string` | 字段注释信息       | 用户ID        |
| MappingType | `string` | 映射的Go类型      | int         |

```gohaml
package model

import (
{{ if $.HasTimestamp }}
    "time"
{{ end }}
)

type {{$.StudlyName}} struct {
{{- range $.Columns }}
	{{.StudlyName}} {{.MappingType}} `json:"{{.Field}}" gorm:"{{ if eq .Field $.PrimaryField }}primaryKey;{{end}}column:{{.Field}}"` // {{.Comment}}
{{- end }}
}

func (*{{$.StudlyName}}) TableName() string {
	return "{{$.TableName}}"
}

func (*{{$.StudlyName}}) PK() string {
	return "{{$.PrimaryField}}"
}
```