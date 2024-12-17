package day6

import (
	"fmt"
	"testing"
)

func BenchmarkAbdi(b *testing.B) {
	f := readFile("input")
	for _, l := range []int{100, 1000, 10000, 100000, 1000000} {
		b.Run(fmt.Sprintf("%d days", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				simulate2(f, i)
			}
		})
	}
}

func BenchmarkDag(b *testing.B) {
	f := readFile("input")
	for _, l := range []int{100, 1000, 10000, 100000, 1000000} {
		b.Run(fmt.Sprintf("%d days", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				part2_no_modulo(f, i)
			}
		})
	}
}
