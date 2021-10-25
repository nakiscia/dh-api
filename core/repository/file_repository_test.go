package repository

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFileRepository(t *testing.T) {
	repository:= NewFileRepository("/filepath")
	assert.NotNil(t, repository)
}

func TestWrite(t *testing.T){

	type testCase struct{
		name string
		testData interface{}
		expectedData interface{}
		makeMock func()
	}

	repository:= NewFileRepository("/filepath")

	testCases := []testCase{
		{
			name:         "Write to file successfully",
			testData: map[string]string{"test":"test"},
			expectedData: nil,
			makeMock: func() {
				writeToFile = func(filePath string, data []byte) error {
					return nil
				}
			},
		},
		{
			name:         "Write to file error on writing data",
			testData: "11",
			expectedData: fmt.Errorf(ErrorOnWritingData, "\"11\""),
			makeMock: func() {
				writeToFile = func(filePath string, data []byte) error {
					return fmt.Errorf(ErrorOnWritingData, "11")
				}
			},
		},
	}

	for _,c := range testCases{
		if c.makeMock != nil {
			c.makeMock()
		}

		err := repository.Write(c.testData)
		assert.Equal(t, c.expectedData,err)
	}
}


func TestRead(t *testing.T){

	type testCase struct{
		name string
		testData interface{}
		expectedData interface{}
		makeMock func()
	}

	repository:= NewFileRepository("/filepath")

	testCases := []testCase{
		{
			name:         "Read file successfully",
			testData: map[string]string{"test":"test"},
			expectedData: `{"test":"test"}`,
			makeMock: func() {
				readFile = func(filePath string) ([]byte, error) {
					return []byte(`{"test":"test"}`),nil
				}
			},
		},
		{
			name:         "Read file error on reading data",
			testData: "",
			expectedData: errors.New(ErrorOnReadingData),
			makeMock: func() {
				readFile = func(filePath string) ([]byte, error) {
					return nil,errors.New(ErrorOnReadingData)
				}
			},
		},
	}

	for _,c := range testCases{
		if c.makeMock != nil {
			c.makeMock()
		}

		data, err := repository.Read()
		if data != nil {
			assert.Equal(t, c.expectedData ,string(data.([]byte)))
		} else {
			assert.Equal(t, c.expectedData,err)
		}

	}
}