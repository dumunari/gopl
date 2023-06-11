package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args[1:] {
		fmt.Printf("Index: %d Value: %s\n", index, value)
	}
}
