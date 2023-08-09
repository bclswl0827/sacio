package sacio

import (
	"bufio"
	"os"
	"reflect"
)

// Read SAC file to structured SACData
func (s *SACData) Read(filename string) error {
	type SACRawData struct {
		BitOrder  int
		DataBytes []byte
	}

	// Open SAC file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read file to bytes
	var bytes []byte
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		bytes = append(bytes, buffer[:n]...)
	}

	// Check file bit order
	sacRawData := &SACRawData{
		BitOrder:  MSBFIRST,
		DataBytes: bytes,
	}
	if assembleInt32(sacRawData.DataBytes[75*4+4:75*4+8], MSBFIRST) != 6 {
		sacRawData.BitOrder = LSBFIRST
	}

	// Read data to struct
	var (
		dataOrder = sacRawData.BitOrder
		dataBytes = sacRawData.DataBytes
	)
	for i, j := 0, 0; i < HEADER_LENGTH; j++ {
		header, err := getVariableByIndex(j)
		if err != nil {
			return err
		}

		var (
			dataVariable = header.Variable
			dataSlice    = dataBytes[i : i+header.Width]
		)

		i += header.Width
		if dataVariable == "UNUSED" || dataVariable == "INTERNAL" {
			continue
		}

		t := reflect.ValueOf(s).Elem()
		switch header.DataType {
		case "F":
			result := assembleFloat32(dataSlice, dataOrder)
			err = setStructField(t, dataVariable, F(result))
		case "N":
			result := assembleInt32(dataSlice, dataOrder)
			err = setStructField(t, dataVariable, N(result))
		case "L":
			result := assembleBool(dataSlice, dataOrder)
			err = setStructField(t, dataVariable, L(result))
		case "K":
			result := assembleString(dataSlice)
			err = setStructField(t, dataVariable, K(result))
		case "I":
			result := assembleEnum(dataSlice, dataOrder)
			err = setStructField(t, dataVariable, I(result))
		}

		if err != nil {
			return err
		}
	}

	// Read body section
	var bodyData []F
	for i := HEADER_LENGTH; i < len(dataBytes); i += 4 {
		bodyData = append(bodyData, assembleFloat32(dataBytes[i:i+4], dataOrder))
	}

	s.Body = bodyData
	return nil
}
