package machine

import (
	"errors"
	"go-orbstack/internal/orbstack"
	"os/exec"
	"strings"
)

func List() ([]*orbstack.Machine, error) {

	cmdExec := exec.Command("sh", "-c", "orb list")

	out, err := cmdExec.Output()

	if err != nil {
		return nil, errors.New(string(out) + " " + err.Error())
	}

	return parseMachines(string(out))

}

func parseMachines(out string) ([]*orbstack.Machine, error) {

	var instances []*orbstack.Machine

	allLines := strings.Split(strings.TrimSpace(out), "\n")

	for _, line := range allLines {
		parts := strings.Fields(line)
		instances = append(instances, &orbstack.Machine{Name: parts[0], State: parts[1], Distro: parts[2], Version: parts[3], Arch: parts[4]})
	}

	return instances, nil
}
