package rest

import (
	"github.com/nathaliaguayos/msg-receiver/internal/handlers/handlersfakes"
	"github.com/rs/zerolog"
	"testing"
)

func TestNewRestClient(t *testing.T) {
	t.Run("should return an error when logger is nil", func(t *testing.T) {
		_, err := NewRestClient(nil, nil, 0)
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("should return an error when jwtHandler is nil", func(t *testing.T) {
		log := zerolog.Nop()
		_, err := NewRestClient(&log, nil, 0)
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("should return a new rest client", func(t *testing.T) {
		log := zerolog.Nop()
		jwtHandler := &handlersfakes.FakeJWTHandler{}
		client, err := NewRestClient(&log, jwtHandler, 0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if client == nil {
			t.Error("expected a client, got nil")
		}
	})
}
