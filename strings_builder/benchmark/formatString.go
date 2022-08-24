package talk

import (
	"strings"
)

func formatString(s []string) (ret string) {
	for _, str := range s {
		ret += str
	}
	return
}

func builderString(s []string) string {
	var sb strings.Builder
	for _, str := range s {
		sb.WriteString(str) // HL
	}
	return sb.String()
}
