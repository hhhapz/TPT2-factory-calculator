package factory

import (
	"encoding/json"
	"io"
	"os"
)

var Data *Factory

func Parse(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &Data)
	return
}
