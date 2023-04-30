package helper

import "testing"

func BenchmarkHelloWorldTableTest(b *testing.B) {
	benchmarks := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:    "HelloWorld(Fransiskus)",
			request: "Fransiskus",
		},
		{
			name:    "HelloWorld(Andika)",
			request: "Andika",
		},
		{
			name:    "HelloWorld(Setiawan)",
			request: "Setiawan",
		},
		{
			name:    "HelloWorld(Dikakoko)",
			request: "Dikakoko",
		},
	}

	for _, test := range benchmarks {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(test.request)
			}
		})
	}
}
