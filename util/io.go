package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadString() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		fmt.Println("\033[31mError\033[0m: " + err.Error())
		os.Exit(0)
	}

	return scanner.Text()
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
