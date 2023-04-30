package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHelloWorldTestify(t *testing.T) {
	result := HelloWorld("Dika")
	assert.Equal(t, "Hello Dika", result, "Result must be 'Hello Dika'")

	// assert.Equal -> Ini ketika unit test passed atau failed program ini akan tetap dijalankan
	fmt.Println("Test Done")
}

func TestHayWorldTestify(t *testing.T) {
	result := HayWorld("Dika")
	require.Equal(t, "Hello Dika", result, "Result must be 'Hello Dika'")

	// require.Equal -> Ini ketika unit test failed program ini tidak akan dijalankan
	fmt.Println("Test Done")
}
