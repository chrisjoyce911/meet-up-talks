package talk

import (
	"testing"
)

var result string

// START OMIT

func data() []string {
	s := []string{"Test String"}
	for i := 0; i < 25; i++ {
		s = append(s, s[len(s)-1]+"Test String")
	}
	return s
}

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
