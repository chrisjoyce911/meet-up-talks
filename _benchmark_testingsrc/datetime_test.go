package validator

import (
	"testing"
)

func benchmarkDateRegex(v string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = DateRegex(v)
	}
}

func BenchmarkDateRegex_1(b *testing.B) { benchmarkDateRegex("", b) }
func BenchmarkDateRegex_2(b *testing.B) { benchmarkDateRegex("     ", b) }
func BenchmarkDateRegex_3(b *testing.B) { benchmarkDateRegex("2020-01-01", b) }
