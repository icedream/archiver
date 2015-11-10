// +build windows

package extractor

import (
	"errors"
	"os/exec"
)

func tarFromShell(src, dest string) (*exec.Cmd, error) {
	return nil, errors.New("tarFromShell function undefined on windows")
}
