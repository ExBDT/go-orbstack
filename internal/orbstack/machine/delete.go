package machine

import (
	"errors"
	"os/exec"
)

func Delete(name string) (string, error) {

	var args = []string{"delete", name, "-f"}

	_, err := exec.Command("orb", args...).CombinedOutput()

	if err != nil {
		return "", errors.New(err.Error())
	}
	return "ok", nil

}
