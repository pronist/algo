package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func Getfloats(fileName string) ([3]float64, error) {
	var data [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		return data, err
	}

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		data[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return data, err
		}
	}
	if err = file.Close(); err != nil {
		return data, err
	}
	if err = scanner.Err(); err != nil {
		return data, err
	}
	return data, nil
}
