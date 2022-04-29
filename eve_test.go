package main

import "testing"

func BenchmarkEveEat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e := &eve{}
		e.eat()
	}
}
