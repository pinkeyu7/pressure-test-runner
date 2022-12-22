package file

import (
	"fmt"
	"os"
	"pressure-test-runner/config"
)

func Logs(fileName, title, content string) error {
	fileName = fmt.Sprintf("%s%s", config.GetOutputFilePath(), fileName)

	var _, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		var file, err = os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	fo, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	_, err = fo.WriteString(title + content)
	if err != nil {
		return err
	}

	return nil
}
