package util

import (
	"encoding/json"
	"io/ioutil"
)

// ReadJSONFile reads a JSON file from the specified path into an object.
func ReadJSONFile(path string, object interface{}) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}

	err = json.Unmarshal([]byte(file), object)
	if err != nil {
		return err
	}

	return nil
}

// WriteJSONFile writes a JSON file into the specified path from an object.
func WriteJSONFile(path string, object interface{}) error {
	file, err := json.MarshalIndent(object, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
