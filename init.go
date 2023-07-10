package sacio

import "reflect"

func (s *SACData) Init() error {
	t := reflect.ValueOf(s).Elem()
	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Type().Field(i).Name
		fieldValue := t.Field(i).Interface()
		switch fieldValue.(type) {
		case F:
			setStructField(t, fieldName, F(-12345.0))
		case N:
			setStructField(t, fieldName, N(-12345))
		case L:
			setStructField(t, fieldName, L(false))
		case I:
			setStructField(t, fieldName, I(""))
		case K:
			setStructField(t, fieldName, K("-12345  "))
		}
	}

	return nil
}
