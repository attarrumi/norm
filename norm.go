package norm

import (
	"math"
	"sort"
)

func MinMaxScale(data []float64) []float64 {
	result := make([]float64, len(data))
	minValue := data[0]
	maxValue := data[0]
	for _, value := range data {
		if value < minValue {
			minValue = value
		}
		if value > maxValue {
			maxValue = value
		}
	}
	rangeValue := maxValue - minValue
	for i, value := range data {
		result[i] = (value - minValue) / rangeValue
	}
	return result
}

func LogScale(data []float64) []float64 {
	result := make([]float64, len(data))
	for i, value := range data {
		result[i] = math.Log(value)
	}
	return result
}

func ZScoreScale(data []float64) []float64 {
	result := make([]float64, len(data))
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	mean := sum / float64(len(data))

	varianceSum := 0.0
	for _, value := range data {
		varianceSum += math.Pow(value-mean, 2)
	}
	stdDev := math.Sqrt(varianceSum / float64(len(data)-1))

	for i, value := range data {
		result[i] = (value - mean) / stdDev
	}
	return result
}

func median(data []float64) float64 {
	n := len(data)
	sorted := make([]float64, n)
	copy(sorted, data)
	sort.Float64s(sorted)
	if n%2 == 0 {
		return (sorted[n/2-1] + sorted[n/2]) / 2
	}
	return sorted[n/2]
}

func RobustScale(data []float64) []float64 {
	result := make([]float64, len(data))
	q2 := median(data)
	q1 := median(data[:len(data)/2])
	q3 := median(data[len(data)/2:])
	iqr := q3 - q1
	for i, value := range data {
		result[i] = (value - q2) / iqr
	}
	return result
}
