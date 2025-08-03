package fileservice

import "os"

type FileService interface {
	CreateDir(name string) error
	CreateFile(filename string) error
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, contents []byte) error
	AppendFile(filename string, contents []byte) error
}

type FileOperations struct{}

func (f FileOperations) CreateDir(name string) error {
	return os.MkdirAll(name, os.ModePerm)
}

func (f FileOperations) CreateFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	return file.Close()
}

func (f FileOperations) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (f FileOperations) WriteFile(filename string, contents []byte) error {
	return os.WriteFile(filename, contents, 0644)
}

func (f FileOperations) AppendFile(filename string, contents []byte) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(contents)
	return err
}
