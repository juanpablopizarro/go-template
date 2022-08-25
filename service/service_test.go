package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidHashInput(t *testing.T) {
	srv := NewService()
	_, err := srv.Hash("")
	assert.NotNil(t, err)
}

func TestValidHash(t *testing.T) {
	srv := NewService()
	h, err := srv.Hash("some_input")
	assert.Nil(t, err)
	assert.NotNil(t, h)
}
