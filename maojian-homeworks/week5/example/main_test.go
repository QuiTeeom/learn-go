package main

import "testing"

func BenchmarkNewLimiter(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Add()
		}
	})
}
