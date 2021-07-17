package datafile

import (
	"bufio"
	"os"
)

func GetStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = file.Close(); err != nil {
		return nil, err
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
