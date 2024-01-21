package util

import (
	"fmt"
	"os"
)

func ReadString() string {
	ln := ""
	fmt.Scanf("%s", &ln)

	return ln
}

func WriteStringToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return nil
}
