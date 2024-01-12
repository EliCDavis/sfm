package sfm

import (
	"encoding/json"
	"io"
	"os"
)

func LoadJSONFile[T any](filename string) (T, error) {
	var data T
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return data, err
	}
	return data, json.Unmarshal(fileData, &data)
}

func ReadJSON[T any](in io.Reader) (T, error) {
	var data T
	fileData, err := io.ReadAll(in)
	if err != nil {
		return data, err
	}
	return data, json.Unmarshal(fileData, &data)
}
