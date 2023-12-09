package bb

import (
	"os"
)

func NewTempFile(data []byte) (string, error) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		return "", err
	}
	path := f.Name()
	if data != nil {
		_, err = f.Write(data)
		if err != nil {
			return "", err
		}
		err = f.Close()
		if err != nil {
			return path, err
		}
	}
	return path, nil
}

func GetTempDir() (string, error) {
	return os.TempDir(), nil
}

func NewTempDir() (string, error) {
	return os.MkdirTemp("", "")
}
