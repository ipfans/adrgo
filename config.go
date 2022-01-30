package adrgo

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Language string
	Path     string
	Digits   int
}

func MarshalConfig(conf Config) []byte {
	b, _ := yaml.Marshal(conf)
	return b
}

func IsInitedConfig() bool {
	return pathExists(".adr.yml")
}

func ReadConfig() (conf Config, err error) {
	var b []byte
	b, err = os.ReadFile(".adr.yml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(b, &conf)
	return
}
