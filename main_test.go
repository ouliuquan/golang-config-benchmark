package main

import (
	"testing"
)

func BenchmarkTomlConfig(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parseToml()
	}
}

func BenchmarkJsonConfig(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parseJson()
	}
}
