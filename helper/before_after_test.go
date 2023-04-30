package helper

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMain(t *testing.M) {
	fmt.Println("Before unit test run")

	t.Run()

	fmt.Println("After unti test run")
}

func TestHelloWorldBeforeAfter(t *testing.T) {
	result := HelloWorld("Dika")
	require.Equal(t, "Hello Dika", result, "Result must be 'Hello Dika'")
}
