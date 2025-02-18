package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func LoadFloatFromFile(fileName string) (float64, error) {
	data, fileReadError := os.ReadFile(fileName)

	if fileReadError != nil {
		return 0, errors.New("failed to read file")
	}

	valueText := string(data)
	value, parsingError := strconv.ParseFloat(valueText, 64)

	if parsingError != nil {
		return 0, errors.New("failed to parse balance value")
	}

	return value, nil
}
