package utils

import (
    "errors"
    "math"
)
func Round(f float64, places int) (float64, error) {
    if places < 0 {
        return 0, errors.New("Places cannot be less than 0")
    }
    shift := math.Pow(10, float64(places))
    return math.Round(f * shift) / shift, nil
}

func CalculateAverage(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("Cannot calculate average of empty array")
	}
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers)), nil
}