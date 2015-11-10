// +build !windows

package extractor

import (
	"os"
	"os/exec"
)

func tarFromShell(src, dest string) (*exec.Cmd, error) {
	tarPath, err := exec.LookPath("tar")
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(dest, 0755)
	if err != nil {
		return nil, err
	}

	return exec.Command(tarPath, "pzxf", src, "-C", dest), nil
}
