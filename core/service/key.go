package service

import (
	"dh-api/core/repository"
	"encoding/json"
	"errors"
	"time"
)

type KeyValueMap map[string]string

const InvalidKeyError = "key must be a string value"
const KeyNotFound = "key not found"

type KeyService struct{
	tickerDuration time.Duration
	fileRepo repository.FileRepositoryInterface
	keysAndValue KeyValueMap
}

type KeyServiceInterface interface{
	SetKey(key, value string) error
	GetKey(key string) (string,error)
}

func NewKeyService(repo repository.FileRepositoryInterface, tickerDuration time.Duration) *KeyService{
	service := &KeyService{
		tickerDuration: tickerDuration,
		fileRepo:       repo,
	}

	read, err := repo.Read()
	if err == nil && read != nil {
		_ = json.Unmarshal(read.([]byte), &service.keysAndValue)
	} else{
		service.keysAndValue = map[string]string{}
	}

	go service.persist()

	return service
}

func (s *KeyService) SetKey(key, value string) error{
	if key == "" {
		return errors.New(InvalidKeyError)
	}

	s.keysAndValue[key] = value
	return nil
}

func (s *KeyService) GetKey(key string) (string,error) {
	if value,ok:= s.keysAndValue[key]; ok {
		return value,nil
	}

	return "", errors.New(KeyNotFound)
}

func (s *KeyService) persist() {
	if s.tickerDuration.Seconds() <= 0 {
		return
	}

	ticker := time.NewTicker((s.tickerDuration) * time.Millisecond)

	for {
		select {
		case <- ticker.C:
			_ = s.fileRepo.Write(s.keysAndValue)
		}
	}
}