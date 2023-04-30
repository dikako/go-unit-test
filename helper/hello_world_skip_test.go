package helper

import (
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestHelloWorldSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test tidak bisa jalan di Mac")
	}

	result := HelloWorld("Dika")
	require.Equal(t, "Hello Dika", result, "Result must be 'Hello Dika'")
}
