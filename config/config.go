package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type modelConfig struct {
	Dirname      string            `yaml:"dirname" json:"dirname"`             // 目录
	Template     string            `yaml:"template" json:"template"`           // 模板地址
	RepoDirname  string            `yaml:"repo_dirname" json:"repo_dirname"`   // repo 目录
	RepoTemplate string            `yaml:"repo_template" json:"repo_template"` // repo 模板地址
	Types        map[string]string `yaml:"types" json:"types"`                 // 类型映射
}

type Config struct {
	// 项目目录
	ProjectPath string      `yaml:"project_path" json:"project_ath"`
	DB          dbConfig    `yaml:"db" json:"db"`
	Model       modelConfig `yaml:"model" json:"model"`
}

func Load(filename string) *Config {

	conf := &Config{}
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("读取配置文件错误: %v", err)
	}

	if yaml.Unmarshal(content, conf) != nil {
		panic(fmt.Sprintf("解析 config.yaml 读取错误: %v", err))
	}

	return conf
}
