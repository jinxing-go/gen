package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type dbConfig struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Name     string `yaml:"name" json:"name"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

func (d *dbConfig) connection() *sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func NewDefaultDB(config *Config) *sql.DB {
	return config.DB.connection()
}
