package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func Getfloats(fileName string) ([]float64, error) {
	var data []float64
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		data = append(data, number)
	}
	if err = file.Close(); err != nil {
		return data, err
	}
	if err = scanner.Err(); err != nil {
		return data, err
	}
	return data, nil
}
