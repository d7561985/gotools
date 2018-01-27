//
package cfg

import (
	"os"
	"fmt"
	"bufio"
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

func GetFileData(file string) (string, error) {
	f, err := os.Open(file)

	if err != nil {
		return "", fmt.Errorf("Can't open settings.json for game %q\n", file)
	}

	defer f.Close()

	input := bufio.NewScanner(f)
	var data = make([]byte, 0, 10000)
	for input.Scan() {
		data = append(data, input.Bytes()...)
	}
	if input.Err() != nil {
		return "", fmt.Errorf("Can't open settings.json for game %q err: %q\n", file, input.Err())
	}

	return string(data), nil
}
