package machine

import (
	"errors"
	"go-orbstack/internal/orbstack"
	"os/exec"
)

type CreateReq struct {
	Name         string
	Architecture string
	Distribution string
	Version      string
	User         string
	Password     string
}

func Create(createReq *CreateReq) (*orbstack.Machine, error) {

	var args = []string{"create"}

	if createReq.Distribution != "" {
		args = append(args, createReq.Distribution)
	}

	if createReq.Name != "" {
		args = append(args, createReq.Name)
	}

	result := exec.Command("orb", args...)
	_, err := result.CombinedOutput()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &orbstack.Machine{Name: createReq.Name, Distro: createReq.Distribution}, nil

}
