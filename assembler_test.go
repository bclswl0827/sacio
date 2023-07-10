package sacio

import (
	"math"
	"reflect"
	"testing"
)

func TestAssembleString(t *testing.T) {
	data := []byte{'H', 'e', 'l', 'l', 'o'}
	expectedResult := K(data)

	result := assembleString(data)
	if result != expectedResult {
		t.Errorf("Expected assembled string: %v, got: %v", expectedResult, result)
	}
}

func TestAssembleBool(t *testing.T) {
	// Test case: LSBFIRST, true
	data := []byte{0x01, 0x00, 0x00, 0x00}
	expectedResult := L(true)

	result := assembleBool(data, LSBFIRST)
	if result != expectedResult {
		t.Errorf("Expected assembled bool for LSBFIRST: %v, got: %v", expectedResult, result)
	}

	// Test case: LSBFIRST, false
	data = []byte{0x00, 0x00, 0x00, 0x00}
	expectedResult = L(false)

	result = assembleBool(data, LSBFIRST)
	if result != expectedResult {
		t.Errorf("Expected assembled bool for LSBFIRST: %v, got: %v", expectedResult, result)
	}

	// Test case: MSBFIRST, true
	data = []byte{0x00, 0x00, 0x00, 0x01}
	expectedResult = L(true)

	result = assembleBool(data, MSBFIRST)
	if result != expectedResult {
		t.Errorf("Expected assembled bool for MSBFIRST: %v, got: %v", expectedResult, result)
	}

	// Test case: MSBFIRST, false
	data = []byte{0x00, 0x00, 0x00, 0x00}
	expectedResult = L(false)

	result = assembleBool(data, MSBFIRST)
	if result != expectedResult {
		t.Errorf("Expected assembled bool for MSBFIRST: %v, got: %v", expectedResult, result)
	}
}

func TestAssembleInt32(t *testing.T) {
	// Test case: LSBFIRST
	data := []byte{0x78, 0x56, 0x34, 0x12}
	expectedResult := N(0x12345678)

	result := assembleInt32(data, LSBFIRST)
	if result != expectedResult {
		t.Errorf("Expected assembled int32 for LSBFIRST: %v, got: %v", expectedResult, result)
	}

	// Test case: MSBFIRST
	data = []byte{0x12, 0x34, 0x56, 0x78}
	expectedResult = N(0x12345678)

	result = assembleInt32(data, MSBFIRST)
	if result != expectedResult {
		t.Errorf("Expected assembled int32 for MSBFIRST: %v, got: %v", expectedResult, result)
	}
}

func TestAssembleFloat32(t *testing.T) {
	// Test case: LSBFIRST
	data := []byte{0x7B, 0x14, 0xAE, 0x47}
	expectedResult := F(math.Float32frombits(0x47AE147B))

	result := assembleFloat32(data, LSBFIRST)
	if result != expectedResult {
		t.Errorf("Expected assembled float32 for LSBFIRST: %v, got: %v", expectedResult, result)
	}

	// Test case: MSBFIRST
	data = []byte{0x47, 0xAE, 0x14, 0x7B}
	expectedResult = F(math.Float32frombits(0x47AE147B))

	result = assembleFloat32(data, MSBFIRST)
	if result != expectedResult {
		t.Errorf("Expected assembled float32 for MSBFIRST: %v, got: %v", expectedResult, result)
	}
}

func TestAssembleEnum(t *testing.T) {
	// Test case: LSBFIRST, matching enum
	data := []byte{0x01, 0x00, 0x00, 0x00}
	expectedResult := I("itime")

	result := assembleEnum(data, LSBFIRST)
	if result != expectedResult {
		t.Errorf("Expected assembled enum for LSBFIRST: %v, got: %v", expectedResult, result)
	}

	// Test case: MSBFIRST, matching enum
	data = []byte{0x00, 0x00, 0x00, 0x01}
	expectedResult = I("itime")

	result = assembleEnum(data, MSBFIRST)
	if result != expectedResult {
		t.Errorf("Expected assembled enum for MSBFIRST: %v, got: %v", expectedResult, result)
	}

	// Test case: LSBFIRST, no matching enum
	data = []byte{0x02, 0x00, 0x00, 0x00}
	expectedResult = I("irlim")

	result = assembleEnum(data, LSBFIRST)
	if result != expectedResult {
		t.Errorf("Expected assembled enum for LSBFIRST: %v, got: %v", expectedResult, result)
	}

	// Test case: MSBFIRST, no matching enum
	data = []byte{0x00, 0x00, 0x00, 0x02}
	expectedResult = I("irlim")

	result = assembleEnum(data, MSBFIRST)
	if result != expectedResult {
		t.Errorf("Expected assembled enum for MSBFIRST: %v, got: %v", expectedResult, result)
	}
}

func TestDisassembleFloat32(t *testing.T) {
	// Test case: LSBFIRST
	data := F(3.14159)
	expectedResult := []byte{0x40, 0x49, 0x0F, 0xD0}

	result := disassembleFloat32(data, LSBFIRST)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected disassembled float32 for LSBFIRST: %v, got: %v", expectedResult, result)
	}

	// Test case: MSBFIRST
	data = F(3.14159)
	expectedResult = []byte{0xD0, 0x0F, 0x49, 0x40}

	result = disassembleFloat32(data, MSBFIRST)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected disassembled float32 for MSBFIRST: %v, got: %v", expectedResult, result)
	}
}
