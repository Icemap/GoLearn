package utils

import (
	"main/model"
	"os"
	"encoding/json"
	"io/ioutil"
)

func ConfUtils() model.Conf {
	var buf []byte
	buf, _ = ReadAll("conf.json")
	var conf model.Conf
	json.Unmarshal(buf, &conf)
	return conf
}

func ReadAll(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}