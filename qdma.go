package main

import (
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		LabelFind(arg) // First pass to find all labels.
		Assemble(arg)  // Second pass to assemble the instructions.
	}
}
