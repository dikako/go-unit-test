package helper

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSubTest(t *testing.T) {
	t.Run("Fransiskus", func(t *testing.T) {
		result := HelloWorld("Fransiskus")
		require.Equal(t, "Hello Fransiskus", result, "Result must be 'Hello Fransiskus'")
	})

	t.Run("Dika", func(t *testing.T) {
		result := HelloWorld("Dika")
		require.Equal(t, "Hello Dika", result, "Result must be 'Hello Dika'")
	})
}
