package check

import (
	"errors"
	"io"
	"os/exec"
)

func Checker(err error) {
	if err != nil {
		if err == io.EOF {
			return
		}
		// https://pkg.go.dev/os/exec?GOOS=windows
		if errors.Is(err, exec.ErrDot) {
			return
		}
		panic(err)
	}
}
