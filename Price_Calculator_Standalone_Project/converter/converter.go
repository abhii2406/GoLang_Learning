package converter

import (
	"errors"
	"strconv"
)

func StringToFloats(strs []string) ([]float64, error) {
		var floats []float64
		for _, strVal := range strs {
			floatval, err := strconv.ParseFloat(strVal,64)
			if err !=nil{
				return nil,errors.New("Failed to convert")
			}
			floats=append(floats, floatval)
		}
		return floats,nil
	}
