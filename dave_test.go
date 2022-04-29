package main

import "testing"

func Test_dave_drink(t *testing.T) {
	t.Run("test drinking", func(t *testing.T) {
		d := &dave{}
		d.drink()
	})
}

func BenchmarkDaveDrink(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := &dave{}
		d.drink()
	}
}
