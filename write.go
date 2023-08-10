package sacio

import (
	"os"
)

func (s *SACData) Write(filePath string, dataBytes []byte) error {
	file, err := os.Create(filePath)
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
