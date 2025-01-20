package logger_test

import (
	"github.com/nathaliaguayos/msg-receiver/config"
	"github.com/nathaliaguayos/msg-receiver/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name   string
		config *config.Config
		assert func(t *testing.T, log *zerolog.Logger)
	}{
		{
			name: "success",
			config: &config.Config{
				ServiceName: "msg-receiver",
			},
			assert: func(t *testing.T, log *zerolog.Logger) {
				assert.NotNil(t, log)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			log := logger.New(tc.config)
			tc.assert(t, log)
		})
	}
}
