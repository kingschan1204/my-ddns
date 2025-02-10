package myconf

import (
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
)

type appConfig struct {
	SecretId  string `yaml:"sid"`
	SecretKey string `yaml:"skey"`
	//主域名
	Domain string `yaml:"domain"`
	// 要更新的域名前缀  如果是顶级域名则是填@
	Target string `yaml:"target"`
}

type AppYaml struct {
	App appConfig `yaml:"app"`
}

var App appConfig

// InitConfig 初始化配置信息
// filePath 配置yaml文件路径
func InitConfig(filePath string) {
	//获取当前目录
	//fmt.Println("1.初始化配置信息.")
	//fmt.Println(os.Getwd())
	//filename := "./config.yaml"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file:%v\n", err)
		return
	}
	y := new(AppYaml)
	yamlFile, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("读取配置文件 config.yaml 失败 %v\n", err)
	}
	err1 := yaml.Unmarshal(yamlFile, y)
	if err1 != nil {
		log.Fatalf("yaml 解码失败: %v\n", err)
	}
	App = y.App
}
