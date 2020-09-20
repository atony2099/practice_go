package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Configs struct {
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Protol    string `yaml:"protol"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	DbName    string `yaml:"dbname"`
	Charset   string `yaml:"charset"`
	MaxIdle   int    `yaml:"maxIdle"`
	MaxActive int    `yaml:"maxActive"`
	ShowSQL   bool   `yaml:"showSQL"`
}

type Config struct {
	MySQLCfg map[string]*Configs `yaml:"mysql"`
}

func main() {

	yamlFile, err := ioutil.ReadFile("./test.yaml")

	fmt.Println(err, "go--")

	var conf Config
	err = yaml.Unmarshal(yamlFile, &conf)

	fmt.Println(err, conf)

}
