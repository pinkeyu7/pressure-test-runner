package file

import (
	"os"
	"path/filepath"
	"pressure-test-runner/config"
)

func Clean() error {
	d, err := os.Open(config.GetOutputFilePath())
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(config.GetOutputFilePath(), name))
		if err != nil {
			return err
		}
	}
	return nil
}
