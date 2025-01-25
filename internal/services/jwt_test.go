package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewJWTService(t *testing.T) {
	jwtService := createJWTService()
	assert.NotNil(t, jwtService)
}

func TestGenerateToken(t *testing.T) {
	jwtService := createJWTService()
	token, err := jwtService.GenerateToken("123")
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestValidateToken(t *testing.T) {
	testCases := []struct {
		name      string
		getToken  func(string, JWTService) string
		assertion func(*testing.T, string, JWTService)
	}{
		{
			name: "Valid token",
			getToken: func(userID string, jwtService JWTService) string {
				token, _ := jwtService.GenerateToken(userID)
				return token
			},
			assertion: func(t *testing.T, token string, jwtService JWTService) {
				assert.NotNil(t, token)
				tok, err := jwtService.ValidateToken(token)
				assert.Nil(t, err)
				assert.NotNil(t, tok)
			},
		}, {
			name: "should success with valid token",
			getToken: func(userID string, jwtService JWTService) string {
				return "invalid"
			},
			assertion: func(t *testing.T, token string, jwtService JWTService) {
				tok, err := jwtService.ValidateToken(token)
				assert.Containsf(t, err.Error(), "token contains an invalid number of segments", "error message: %s", err.Error())
				assert.Nil(t, tok)
			},
		}, {
			name: "should fail with invalid token",
			getToken: func(userID string, jwtService JWTService) string {
				return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJUNUtNbHBiSGpKQ2RQSUtmdFZ5SUJBem5IUEllcThyMCJ9.EDZ45MU8V6tlEvAv1KAZeLtAwRSJgSg2bo5VzwNzdRE"
			},
			assertion: func(t *testing.T, token string, jwtService JWTService) {
				_, err := jwtService.ValidateToken(token)
				assert.Containsf(t, err.Error(), "signature is invalid", "error message: %s", err.Error())
			},
		}, {
			name: "should fail with token containing 'bearer '",
			getToken: func(userID string, jwtService JWTService) string {
				return "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJUNUtNbHBiSGpKQ2RQSUtmdFZ5SUJBem5IUEllcThyMCJ9.EDZ45MU8V6tlEvAv1KAZeLtAwRSJgSg2bo5VzwNzdRE"
			},
			assertion: func(t *testing.T, token string, jwtService JWTService) {
				_, err := jwtService.ValidateToken(token)
				assert.Containsf(t, err.Error(), "tokenstring should not contain 'bearer '", "error message: %s", err.Error())
			},
		}, {
			name: "should fail with malformed token",
			getToken: func(userID string, jwtService JWTService) string {
				return "eyJhbGciOiJIUzI1NiIsInR5cCI6.eyJpc3MiOiJUNUtNbHBiSGpKQ2RQSUtmdFZ5SUJBem5IUEllcThyMCJ9.EDZ45MU8V6tlEvAv1KAZeLtAwRSJgSg2bo5VzwNzdRE"
			},
			assertion: func(t *testing.T, token string, jwtService JWTService) {
				_, err := jwtService.ValidateToken(token)
				assert.Containsf(t, err.Error(), "unexpected end of JSON input", "error message: %s", err.Error())
			},
		}, {
			name: "should fail with a RSA signed token",
			getToken: func(userID string, jwtService JWTService) string {
				return "eyJhbGciOiJSUzI1NiIsImtpZCI6IjEyMyIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0LXVzZXIiLCJpYXQiOjE2Nzg5ODk2MDAsImV4cCI6MTY3OTAwMDQwMH0.KXBqbbZh0OwF7C6yNirIgbrtFStzFgw68M-hkJl18QlHr1TZ3wiVGQkT5O3IQbSR2lUZkgRruZOEas3qFbHMOcIuSvWRlEOHH8Xt-N34XYljA8f0nxWQxrGj67H4QFtmNqN9hIAT9RdoDrMrFukIT6Fnbg5Z22Ff6qzxN3ViD9BLYQnRA9rLZg3T7zL4mhmLQDP_lUeSn1ANSH4DY5-4ESGRQuNkhtIDT8NAdHGWEN2ak-R8jFei3PoDbExZnX8wYXVPSIjfN6tfK9nvawcnxPba9lsOBQzYfwY0Ew6uG5roOGGRoyQ6WBASszGpW7sEr82bBPl6l9o2HnDhW1ZMDqzw"
			},
			assertion: func(t *testing.T, token string, jwtService JWTService) {
				_, err := jwtService.ValidateToken(token)
				assert.Containsf(t, err.Error(), "signature is invalid", "error message: %s", err.Error())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jwtService := createJWTService()
			token := tc.getToken("123", jwtService)
			tc.assertion(t, token, jwtService)
		})
	}
}

func createJWTService() JWTService {
	secretKey := "secret"
	issuer := "issuer"
	jwtService := NewJWTService(secretKey, issuer)
	return jwtService
}
