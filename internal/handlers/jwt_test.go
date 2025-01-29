package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nathaliaguayos/msg-receiver/internal/services/servicesfakes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewJWTHandler(t *testing.T) {
	jwtFakeService := &servicesfakes.FakeJWTService{}
	handler := NewJWTHandler(jwtFakeService)
	assert.NotNil(t, handler)
}

func TestGenerateToken(t *testing.T) {
	testCases := []struct {
		name               string
		requestBody        map[string]string
		jwtService         *servicesfakes.FakeJWTService
		expectedStatusCode int
		assert             func(t *testing.T, w *httptest.ResponseRecorder)
	}{
		{
			name: "should success and retrieve token",
			requestBody: map[string]string{
				"user_id": "1",
			},
			jwtService: &servicesfakes.FakeJWTService{
				GenerateTokenStub: func(userID string) (string, error) {
					return "token", nil
				},
			},
			expectedStatusCode: http.StatusOK,
			assert: func(t *testing.T, w *httptest.ResponseRecorder) {
				var response map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.Nil(t, err)
				assert.Equal(t, "token", response["token"])
			},
		}, {
			name:               "Should return status code 400 when request body is empty",
			requestBody:        map[string]string{},
			jwtService:         &servicesfakes.FakeJWTService{},
			expectedStatusCode: http.StatusBadRequest,
			assert: func(t *testing.T, w *httptest.ResponseRecorder) {
				var response map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.Nil(t, err)
				assert.Containsf(t, response, "error", "response should contain error message")

			},
		}, {
			name: "Should return status code 500 when jwt service fails",
			requestBody: map[string]string{
				"user_id": "1",
			},
			jwtService: &servicesfakes.FakeJWTService{
				GenerateTokenStub: func(userID string) (string, error) {
					return "", assert.AnError
				},
			},
			expectedStatusCode: http.StatusInternalServerError,
			assert: func(t *testing.T, w *httptest.ResponseRecorder) {
				var response map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.Nil(t, err)
				assert.Containsf(t, response, "error", "response should contain error message")
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			gin.Default()
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			jsonValue, _ := json.Marshal(tc.requestBody)
			c.Request = httptest.NewRequest(http.MethodPost, "/v1/token", bytes.NewBuffer(jsonValue))
			c.Request.Header.Set("Content-Type", "application/json")

			handler := NewJWTHandler(tc.jwtService)

			handler.GenerateToken(c)
			assert.Equal(t, tc.expectedStatusCode, w.Code)
			tc.assert(t, w)
		})
	}
}
