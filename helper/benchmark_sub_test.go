package helper

import "testing"

func BenchmarkHelloWorldSubTest(b *testing.B) {
	b.Run("Fransiskus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fransiskus")
		}
	})

	b.Run("Dika", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Dika")
		}
	})
}
