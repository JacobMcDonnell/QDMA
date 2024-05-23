package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// Map for labels and their addresses
var labels map[string]uint = make(map[string]uint)

func LabelFind(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input := bufio.NewScanner(file)

	for i := 0; input.Scan(); i += 4 {
		k := input.Text()
		if k == "" {
			i -= 4
			continue
		} else if k == ".text" || k == ".data" {
			i += 4
			continue
		}
		s, err := Parse(k)
		if err != nil {
			panic(err)
		}

		if len(s) == 1 && s[0] == "" {
			i -= 4
			continue
		}

		hasLabel, err := regexp.MatchString("^.*:", s[0])
		if err != nil {
			panic(err)
		}
		if hasLabel {
			labels[strings.ReplaceAll(s[0], ":", "")] = uint(i)
			if len(s) == 1 {
				i -= 4
			}
		}
	}
}
