package myconf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type appConfig struct {
	SecretId  string `yaml:"sid"`
	SecretKey string `yaml:"skey"`
	//主域名
	Domain string `yaml:"domain"`
}

type appYaml struct {
	App appConfig `yaml:"app"`
}

var App appConfig

func init() {
	//获取当前目录
	//fmt.Println("1.初始化配置信息.")
	fmt.Println(os.Getwd())
	filename := "./config.yaml"
	y := new(appYaml)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("读取配置文件 myconf.yaml 失败 %v\n", err)
	}
	err1 := yaml.Unmarshal(yamlFile, y)
	if err1 != nil {
		log.Fatalf("yaml 解码失败: %v\n", err)
	}
	App = y.App
}
