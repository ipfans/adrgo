package adrgo

import (
	"fmt"
	"io/ioutil"
)

func Init(conf Config) (err error) {
	if IsInitedConfig() {
		return fmt.Errorf(".adr.yml is already inited")
	}
	b := MarshalConfig(conf)
	err = ioutil.WriteFile(".adr.yml", b, 0644)
	if err != nil {
		err = fmt.Errorf("write .adr.yml failed: %w", err)
	}
	return
}
