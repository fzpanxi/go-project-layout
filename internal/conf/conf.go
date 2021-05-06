package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Conf struct {
	Data struct {
		Database struct {
			Driver  string
			Dsn     string
			MaxIdle int
			MaxOpen int
		}
	}
}

func LoadConfig(filename string) (*Conf, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	bc := new(Conf)
	err = yaml.Unmarshal([]byte(data), bc)
	if err != nil {
		return nil, err
	}
	return bc, nil
}
