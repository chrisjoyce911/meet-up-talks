package talk

import (
	"testing"
)

var result string

func data() []string {
	var s []string
	last := "a"
	for i := 0; i < 1000; i++ {
		s = append(s, last)
		last += "a"
	}
	return s
}

// START OMIT
func Benchmark_formatString(b *testing.B) {

	var s string
	for i := 0; i < b.N; i++ {
		s = formatString(data())
	}
	result = s
}

func Benchmark_builderString(b *testing.B) {

	var s string
	for i := 0; i < b.N; i++ {
		s = builderString(data())
	}
	result = s
}

//END OMIT
