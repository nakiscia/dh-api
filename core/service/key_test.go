package service

import (
	"dh-api/core/repository"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewKeyService(t *testing.T) {
	service := NewKeyService(&repository.FileRepository{}, 0)

	assert.NotNil(t, service)
}

func TestSetKey(t *testing.T){
	service := NewKeyService(&repository.FileRepository{}, 0)

	type testCase struct{
		testName string
		expectedValue interface{}
		keyTestData string
		valueTestData string
	}

	testCases := []testCase{
		{
			testName:      "Set key successfully",
			expectedValue: nil,
			keyTestData:   "testKey",
			valueTestData: "testValue",
		},
		{
			testName:      "Set key error on empty key",
			expectedValue: errors.New(InvalidKeyError),
			keyTestData:   "",
			valueTestData: "val",
		},
	}
	
	for _,c := range testCases {
		err := service.SetKey(c.keyTestData,c.valueTestData)
		assert.Equal(t, c.expectedValue, err)
	}
}

func TestGetKey(t *testing.T){
	service := NewKeyService(&repository.FileRepository{}, 0)
	service.keysAndValue = map[string]string{"key":"value"}

	type testCase struct{
		testName string
		expectedValue string
		keyTestData string
		valueTestData string
		expectedError error
	}

	testCases := []testCase{
		{
			testName:      "Get key successfully",
			expectedValue: "value",
			keyTestData:   "key",
			valueTestData: "value",
			expectedError: nil,
		},
		{
			testName:      "Get key error on non-exist key",
			expectedValue: "",
			keyTestData:   "",
			valueTestData: "val",
			expectedError: errors.New(KeyNotFound),
		},
	}

	for _,c := range testCases {
		key, err := service.GetKey(c.keyTestData)
		assert.Equal(t, c.expectedValue, key)
		assert.Equal(t, c.expectedError, err)
	}
}