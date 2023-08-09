package sacio

import "reflect"

func (s *SACData) Encode(bitOrder int) ([]byte, error) {
	buf := make([]byte, HEADER_LENGTH+len(s.Body)*4)
	values := reflect.ValueOf(s).Elem()

	// Go through HEADERS to match variable
	for i, j := 0, 0; i < HEADER_LENGTH; j++ {
		header, err := getVariableByIndex(j)
		if err != nil {
			return nil, err
		}

		result, err := getStructField(values, header.Variable)
		if err != nil {
			return nil, err
		}

		var (
			head = i
			tail = i + header.Width
		)
		i += header.Width

		switch header.DataType {
		case "F":
			copy(buf[head:tail], disassembleFloat32(result.(F), bitOrder))
		case "N":
			copy(buf[head:tail], disassembleInt32(result.(N), bitOrder))
		case "L":
			copy(buf[head:tail], disassembleBool(result.(L), bitOrder))
		case "I":
			copy(buf[head:tail], disassembleEnum(result.(I), bitOrder))
		case "K":
			copy(buf[head:tail], disassembleString(result.(K)))
		}
	}

	// Copy SAC body section
	for i := 0; i < len(s.Body); i++ {
		copy(buf[HEADER_LENGTH+i*4:HEADER_LENGTH+(i+1)*4], disassembleFloat32(s.Body[i], bitOrder))
	}

	return buf, nil
}
