package sacio

import (
	"fmt"
	"reflect"
	"time"
)

// getVariableByIndex returns header variable by index
func getVariableByIndex(index int) (*sacHeader, error) {
	if index > len(headers)-1 {
		return nil, fmt.Errorf("index %d out of range", index)
	}

	return &headers[index], nil
}

// getParsedTimeField returns parsed time field
func getParsedTimeField(t time.Time) (*sacTime, error) {
	var timeField sacTime

	// Fill basic time fields
	timeField.Year = t.Year()
	timeField.Hour = t.Hour()
	timeField.Min = t.Minute()
	timeField.Sec = t.Second()
	timeField.Msec = t.Nanosecond() / 1000000

	// Calculate day of given year
	startOfYear := time.Date(timeField.Year, time.January, 1, 0, 0, 0, 0, time.UTC)
	timeField.Days = t.YearDay() - startOfYear.YearDay() + 1

	return &timeField, nil
}

// getFloat32MaxValue returns the maximum value of a slice of float32
func getFloat32MaxValue(arr []float32) float32 {
	if len(arr) == 0 {
		return 0.0
	}

	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	return max
}

// getFloat32MinValue returns the minimum value of a slice of float32
func getFloat32MinValue(arr []float32) float32 {
	if len(arr) == 0 {
		return 0.0
	}

	min := arr[0]
	for _, num := range arr {
		if num < min {
			min = num
		}
	}

	return min
}

// getFloat32MeanValue returns the average value of a slice of float32
func getFloat32MeanValue(arr []float32) float32 {
	if len(arr) == 0 {
		return 0.0
	}

	sum := float32(0.0)
	for _, num := range arr {
		sum += num
	}

	average := sum / float32(len(arr))
	return average
}

// getStructField gets the value of a struct field with reflection
func getStructField(v reflect.Value, fieldName string) (any, error) {
	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, fmt.Errorf("field %s does not exist", fieldName)
	}

	return field.Interface(), nil
}

// setStructField sets the value of a struct field with reflection
func setStructField(v reflect.Value, fieldName string, fieldValue any) error {
	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("field %s does not exist", fieldName)
	}

	if !field.CanSet() {
		return fmt.Errorf("cannot set field %s", fieldName)
	}

	if field.Type().Kind() != reflect.TypeOf(fieldValue).Kind() {
		return fmt.Errorf("type mismatch for field %s", fieldName)
	}

	field.Set(reflect.ValueOf(fieldValue))
	return nil
}
