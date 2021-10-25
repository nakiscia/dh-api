package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const ErrorOnMarshaling = "error on marshling the %v"
const ErrorOnWritingData = "error on writing data %s"
const ErrorOnReadingData = "error on reading data"

type FileRepository struct {
	persistenceFilePath string
}

type FileRepositoryInterface interface{
	Write(interface{}) error
	Read() (interface{},error)
}

func NewFileRepository(filePath string) FileRepositoryInterface{
	return &FileRepository{persistenceFilePath: filePath}
}

func (r *FileRepository) Write(data interface{}) error {
	if data == nil {
		return nil
	}

	marshal, err := json.Marshal(data)

	if err != nil {
		return fmt.Errorf(ErrorOnMarshaling, data)
	}

	err = writeToFile(r.persistenceFilePath, marshal)

	if err != nil {
		return fmt.Errorf(ErrorOnWritingData, string(marshal))
	}

	return nil
}

var writeToFile = func(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}

func (r *FileRepository) Read() (interface{},error) {
	fileData, err := readFile(r.persistenceFilePath)

	if err != nil {
		return nil, errors.New(ErrorOnReadingData)
	}

	return fileData, nil
}

var readFile = func(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}