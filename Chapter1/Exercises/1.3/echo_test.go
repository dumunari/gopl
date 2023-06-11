package echo_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var args = []string{"1", "2", "3", "4", "5", "6"}

func BenchmarkInefficient(b *testing.B) {
	os.Stdout = nil
	for index, value := range args {
		fmt.Printf("Index: %d Value: %s\n", index, value)
	}
}

func BenchmarkEfficient(b *testing.B) {
	strings.Join(args, " ")
}
