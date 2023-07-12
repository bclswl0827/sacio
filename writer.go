package sacio

import (
	"os"
)

func (s *SACData) Write(filename string, dataBytes []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(dataBytes)
	if err != nil {
		return err
	}

	return nil
}
