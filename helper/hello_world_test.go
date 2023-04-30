package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Dika")
	if result != "Hello Dika" {
		//t.Fail() -> Use this when you not need print error message
		t.Error("Result must be 'Hello Dika'")
	}

	// t.Fail() -> Ini ketika unit test passed atau failed program ini akan tetap dijalankan
	// t.Error(args...) -> Ini ketika unit test passed atau failed program ini akan tetap dijalankan
	fmt.Println("Test Done")
}

func TestHayWorld(t *testing.T) {
	result := HayWorld("Dika")
	if result != "Hay Dika" {
		//t.FailNow() -> Use this when you not need print error message
		t.Fatal("Result must be 'Hay Dika'")
	}

	// t.FailNow() -> Ini ketika unit test failed program ini tidak akan dijalankan
	// t.Fatal(args...) -> Ini ketika unit test failed program ini tidak akan dijalankan
	fmt.Println("Test Done")
}
