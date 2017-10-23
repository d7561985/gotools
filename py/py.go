//
package py

import (
	"os/exec"
)

type Py struct {

}

func (py *Py) Run(code string) ([]byte, error) {
	cmd := exec.Command("python", "-c", code)

	res, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return res, err
}
