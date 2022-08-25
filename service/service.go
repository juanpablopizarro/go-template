package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
)

type Service interface {
	Hash(string) (string, error)
}

type service struct {
}

var (
	ErrInvalidInput = errors.New("invalid input")
)

func (service) Hash(s string) (string, error) {
	if len(strings.TrimSpace(s)) > 0 {
		return getMD5Hash(s), nil
	}
	return "", ErrInvalidInput
}

func NewService() Service {
	return &service{}
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
