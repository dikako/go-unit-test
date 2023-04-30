package helper

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Fransiskus)",
			request:  "Fransiskus",
			expected: "Hello Fransiskus",
		},
		{
			name:     "HelloWorld(Andika)",
			request:  "Andika",
			expected: "Hello Andika",
		},
		{
			name:     "HelloWorld(Setiawan)",
			request:  "Setiawan",
			expected: "Hello Setiawan",
		},
		{
			name:     "HelloWorld(Dikakoko)",
			request:  "Dikakoko",
			expected: "Hello Dikakoko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be 'Hello Fransiskus'")
		})
	}
}
