package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		inputHeaders  http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name: "Happy-path",
			inputHeaders: http.Header{
				"Authorization": []string{"ApiKey secret-123"},
			},
			expectedKey:   "secret-123",
			expectedError: nil,
		},
		{
			name: "No Key",
			inputHeaders: http.Header{
				"Authorization": []string{""},
			},
			expectedKey:   "secret-123",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name:          "No Header",
			inputHeaders:  http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.inputHeaders)
			if (err != nil) != (tt.expectedError != nil) {
				t.Errorf("error: %v, wantErr: %v", err, tt.expectedError)
				return
			}
			if err == nil && got != tt.expectedKey {
				t.Errorf("GetAPIKey() = %v, expected: %v", got, tt.expectedKey)
			}
		})
	}
}
