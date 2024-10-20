package utils

import (
	"io"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return b, file.Close()
}

func WriteFile(path string, payload []byte) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	if err = file.Truncate(0); err != nil {
		return err
	}

	if _, err = file.Seek(0, 0); err != nil {
		return err
	}

	// Write bytes to the file
	_, err = file.Write(payload)
	if err != nil {
		return err
	}

	return nil
}
