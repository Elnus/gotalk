package conf

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func init() {
	ReadConf()
}

type Config struct {
	BceAPP BceAPP `yaml:"bceAPP"`
}
type BceAPP struct {
	Name         string `yaml:"name"`
	ClientID     string `yaml:"clientID"`
	ClientSecret string `yaml:"clientSecret"`
}

var Configs Config

func ReadConf() {
	data, err := os.ReadFile("/etc/gotalk/config.yaml")
	if err != nil {
		log.Println(err)
		return
	}
	err = yaml.Unmarshal(data, &Configs)
	if err != nil {
		log.Println(err)
	}
}
