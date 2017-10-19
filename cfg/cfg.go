//
package cfg

import (
	"os"
	"encoding/json"
)

// read file and decode from json to specified interface. it's can be useful as stand alone interface for future use.
func LoadCfg(s string, i interface{}) error {
	file, err := os.Open(s)
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&i)
	if err != nil {
		return err
	}

	return nil
}
