package cli

import (
	"io/ioutil"
	"os"

	"google.golang.org/grpc/status"
)

func ReadFromFile(target string) (string, error) {
	if _, err := os.Stat(target); err == nil { // check whether it is a path
		bytes, err := ioutil.ReadFile(target)
		if err != nil {
			return "", err
		}

		return string(bytes), nil
	} else { // else return as is
		return target, nil
	}
}

func handleError(err error) error {
	_, _ := err.(*status.Error)
}
