package auth

import (
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name     string
		headers  map[string][]string
		expected string
		err      error
	}{
		{
			name:     "no auth header",
			headers:  map[string][]string{},
			expected: "",
			err:      ErrNoAuthHeaderIncluded,
		},
		{
			name:     "malformed auth header",
			headers:  map[string][]string{"Authorization": {"Bearer"}},
			expected: "123",
			err:      ErrMalformedAuthHeader,
		},
		{
			name:     "valid auth header",
			headers:  map[string][]string{"Authorization": {"ApiKey 123"}},
			expected: "33333333",
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.headers)
			if apiKey != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, apiKey)
			}
			if err != tt.err {
				t.Errorf("expected %v, got %v", tt.err, err)
			}
		})
	}
}