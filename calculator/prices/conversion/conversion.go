package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(data []string) ([]float64, error) {
	var floats []float64
	for _, line := range data {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, errors.New("fail loading")
		}
		floats = append(floats, floatPrice)
	}
	return floats, nil
}
