package config_test

import (
	"github.com/nathaliaguayos/msg-receiver/config"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	initialEnv := getEnv()
	test := []struct {
		it     string
		envs   func(t *testing.T) map[string]string
		assert func(t *testing.T, c *config.Config, err error)
	}{
		{
			it: "valid env vars should return hydrated config",
			envs: func(_ *testing.T) map[string]string {
				return map[string]string{
					config.EnvPrefix + "_SERVICE_NAME": "msg-receiver",
					config.EnvPrefix + "_LOG_LEVEL":    "info",
					config.EnvPrefix + "_SECRET_KEY":   "secret",
					config.EnvPrefix + "_ISSUER":       "userName",
				}
			},
			assert: func(t *testing.T, c *config.Config, err error) {
				require.NoError(t, err, "no error should be returned")
				require.NotNil(t, c, "config should not be nil on success")
				require.Equal(t, &config.Config{
					ServiceName: "msg-receiver",
					LogLevel:    "info",
					SecretKey:   "secret",
					Issuer:      "userName",
					Port:        8080,
					Host:        "0.0.0.0",
					RateLimit:   5,
				}, c, "invalid config returned")
			},
		}, {
			it: "missing required env vars should return error",
			envs: func(_ *testing.T) map[string]string {
				return map[string]string{}
			},
			assert: func(t *testing.T, c *config.Config, err error) {
				require.Error(t, err, "error should be returned on missing required env vars")
				require.Nil(t, c, "config should be nil on error")
			},
		},
	}
	for _, tt := range test {
		for k, v := range tt.envs(t) {
			require.NoError(t, os.Setenv(k, v), "cannot set env var")
		}
		t.Run(tt.it, func(t *testing.T) {
			cfg, err := config.Get()
			tt.assert(t, cfg, err)
		})
		resetEnv(t, initialEnv)
	}
}

func TestNew(t *testing.T) {
	initialEnv := getEnv()
	test := []struct {
		it     string
		envs   func(t *testing.T) map[string]string
		assert func(t *testing.T, c *config.Config, err error)
	}{
		{
			it: "valid env vars should return hydrated config",
			envs: func(_ *testing.T) map[string]string {
				return map[string]string{
					config.EnvPrefix + "_SERVICE_NAME": "msg-receiver",
					config.EnvPrefix + "_LOG_LEVEL":    "info",
					config.EnvPrefix + "_SECRET_KEY":   "secret",
					config.EnvPrefix + "_ISSUER":       "userName",
				}
			},
			assert: func(t *testing.T, c *config.Config, err error) {
				require.NoError(t, err, "no error should be returned")
				require.NotNil(t, c, "config should not be nil on success")
				require.Equal(t, &config.Config{
					ServiceName: "msg-receiver",
					LogLevel:    "info",
					SecretKey:   "secret",
					Issuer:      "userName",
					Port:        8080,
					Host:        "0.0.0.0",
					RateLimit:   5,
				}, c, "invalid config returned")
			},
		},
	}
	for _, tt := range test {
		for k, v := range tt.envs(t) {
			require.NoError(t, os.Setenv(k, v), "cannot set env var")
		}
		t.Run(tt.it, func(t *testing.T) {
			cfg := config.New()
			require.NotNil(t, cfg)
		})
		resetEnv(t, initialEnv)
	}
}

func getEnv() map[string]string {
	envs := make(map[string]string)
	for _, e := range os.Environ() {
		kv := strings.SplitN(e, "=", 2)
		envs[kv[0]] = kv[1]
	}
	return envs
}

func resetEnv(t *testing.T, initialEnv map[string]string) {
	for k := range getEnv() {
		_, present := initialEnv[k]
		if !present {
			require.NoError(t, os.Unsetenv(k))
		}
	}
	for k, v := range initialEnv {
		require.NoError(t, os.Setenv(k, v))
	}
}
