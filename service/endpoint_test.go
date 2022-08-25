package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHashEndpoint(t *testing.T) {
	srv := NewService()
	t.Run("valid input", func(t *testing.T) {
		endpoint := makeHashEndpoint(srv)
		req := HashRequest{"some_input"}
		_, err := endpoint(context.Background(), req)
		assert.Nil(t, err)
	})

	t.Run("invalid input", func(t *testing.T) {
		endpoint := makeHashEndpoint(srv)
		req := HashRequest{""}
		_, err := endpoint(context.Background(), req)
		assert.NotNil(t, err)
		assert.EqualError(t, err, ErrInvalidInput.Error())
	})
}
