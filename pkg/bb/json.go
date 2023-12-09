package bb

import (
	"io"
	"os"

	"github.com/mitchellh/mapstructure"
)

func ReadJsonFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}

func ReadJsonFileAsMap(path string) (map[string]interface{}, error) {
	b, err := ReadJsonFile(path)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	err = mapstructure.Decode(b, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
