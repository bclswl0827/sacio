package sacio

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestGetVariableByIndex(t *testing.T) {
	// Test case: valid index
	index := 0
	expectedHeader := &headers[index]
	header, err := getVariableByIndex(index)
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	if header != expectedHeader {
		t.Errorf("Expected header: %v, got: %v", expectedHeader, header)
	}

	// Test case: out of range index
	index = 1000
	_, err = getVariableByIndex(index)
	expectedError := fmt.Errorf("index %d out of range", index)
	if err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %s, got: %s", expectedError.Error(), err)
	}
}

func TestGetParsedTimeField(t *testing.T) {
	// Test case: valid time
	inputTime := time.Now()
	expectedTimeField := &sacTime{
		Year: inputTime.Year(),
		Hour: inputTime.Hour(),
		Min:  inputTime.Minute(),
		Sec:  inputTime.Second(),
		Msec: inputTime.Nanosecond() / 1000000,
		Days: inputTime.YearDay(),
	}

	timeField, err := getParsedTimeField(inputTime)
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	if !reflect.DeepEqual(timeField, expectedTimeField) {
		t.Errorf("Expected time field: %v, got: %v", expectedTimeField, timeField)
	}
}

func TestGetFloat32MaxValue(t *testing.T) {
	// Test case: non-empty array
	arr := []float32{1.5, 2.8, -3.2, 4.7}
	expectedMax := float32(4.7)

	max := getFloat32MaxValue(arr)
	if max != expectedMax {
		t.Errorf("Expected max value: %f, got: %f", expectedMax, max)
	}

	// Test case: empty array
	emptyArr := []float32{}
	expectedEmptyMax := float32(0.0)

	emptyMax := getFloat32MaxValue(emptyArr)
	if emptyMax != expectedEmptyMax {
		t.Errorf("Expected max value for empty array: %f, got: %f", expectedEmptyMax, emptyMax)
	}
}

func TestGetFloat32MinValue(t *testing.T) {
	// Test case: non-empty array
	arr := []float32{1.5, 2.8, -3.2, 4.7}
	expectedMin := float32(-3.2)

	min := getFloat32MinValue(arr)
	if min != expectedMin {
		t.Errorf("Expected min value: %f, got: %f", expectedMin, min)
	}

	// Test case: empty array
	emptyArr := []float32{}
	expectedEmptyMin := float32(0.0)

	emptyMin := getFloat32MinValue(emptyArr)
	if emptyMin != expectedEmptyMin {
		t.Errorf("Expected min value for empty array: %f, got: %f", expectedEmptyMin, emptyMin)
	}
}

func TestGetFloat32MeanValue(t *testing.T) {
	// Test case: non-empty array
	arr := []float32{1.5, 2.8, -3.2, 4.7}
	expectedMean := float32(1.45)

	mean := getFloat32MeanValue(arr)
	if mean != expectedMean {
		t.Errorf("Expected mean value: %f, got: %f", expectedMean, mean)
	}

	// Test case: empty array
	emptyArr := []float32{}
	expectedEmptyMean := float32(0.0)

	emptyMean := getFloat32MeanValue(emptyArr)
	if emptyMean != expectedEmptyMean {
		t.Errorf("Expected mean value for empty array: %f, got: %f", expectedEmptyMean, emptyMean)
	}
}

func TestGetStructField(t *testing.T) {
	// Test case: valid struct and field name
	type TestStruct struct {
		Foo string
		Bar int
	}

	testStruct := TestStruct{
		Foo: "Hello",
		Bar: 42,
	}

	expectedFieldValue := "Hello"
	fieldValue, err := getStructField(reflect.ValueOf(testStruct), "Foo")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	if fieldValue != expectedFieldValue {
		t.Errorf("Expected field value: %v, got: %v", expectedFieldValue, fieldValue)
	}

	// Test case: invalid field name
	_, err = getStructField(reflect.ValueOf(testStruct), "Baz")
	expectedError := fmt.Errorf("field Baz does not exist")
	if err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %s, got: %s", expectedError.Error(), err)
	}
}

func TestSetStructField(t *testing.T) {
	// Test case: valid struct, field name, and field value
	type TestStruct struct {
		Foo string
		Bar int
	}

	testStruct := &TestStruct{
		Foo: "Hello",
		Bar: 42,
	}

	expectedFieldValue := "World"
	err := setStructField(reflect.ValueOf(testStruct).Elem(), "Foo", expectedFieldValue)
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	if testStruct.Foo != expectedFieldValue {
		t.Errorf("Expected field value: %v, got: %v", expectedFieldValue, testStruct.Foo)
	}

	// Test case: invalid field name
	err = setStructField(reflect.ValueOf(testStruct).Elem(), "Baz", "Test")
	expectedError := fmt.Errorf("field Baz does not exist")
	if err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %s, got: %s", expectedError.Error(), err)
	}

	// Test case: type mismatch
	err = setStructField(reflect.ValueOf(testStruct).Elem(), "Bar", "Invalid")
	expectedTypeError := fmt.Errorf("type mismatch for field Bar")
	if err == nil || err.Error() != expectedTypeError.Error() {
		t.Errorf("Expected error: %s, got: %s", expectedTypeError.Error(), err)
	}
}
