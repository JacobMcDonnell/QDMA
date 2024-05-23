package main

import (
	"bufio"
	"encoding/binary"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type PosPair struct {
	pos     uint
	section uint8
}

// Map for labels and their addresses
var labels map[string]PosPair = make(map[string]PosPair)

var SectionPos [4]uint

const (
	TEXT   = 0
	DATA   = 1
	RODATA = 2
	BSS    = 3
)

func LabelFind(path string, dTemp, rTemp *os.File) {
	// This will track the current section, default is .text
	var Section uint8 = TEXT

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input := bufio.NewScanner(file)
	for input.Scan() {
		i := &SectionPos[Section]
		switch input.Text() {
		case "":
			continue
		case ".text":
			Section = TEXT
			continue
		case ".data":
			Section = DATA
			continue
		case ".rodata":
			Section = RODATA
			continue
		case ".bss":
			Section = BSS
			continue
		}

		s, err := Parse(input.Text())
		if err != nil {
			panic(err)
		}

		if len(s) == 1 && s[0] == "" {
			continue
		}

		hasLabel, err := regexp.MatchString("^.*:", s[0])
		if err != nil {
			panic(err)
		}
		if hasLabel {
			labels[strings.ReplaceAll(s[0], ":", "")] = PosPair{*i, Section}
			if len(s) == 1 {
				continue
			} else {
				s = s[1:]
			}
		}

		offset, err := CalcOffset(s, Section, dTemp, rTemp)
		if err != nil {
			panic(err)
		}
		*i += offset
	}
}

func CalcOffset(s []string, Section uint8, dTemp, rTemp *os.File) (uint, error) {
	var offset uint
	if Section != TEXT {
		b, err := EncodeData(s)
		if err != nil {
			return 0, err
		}
		if Section == DATA {
			_, err := dTemp.Write(b)
			if err != nil {
				return 0, err
			}
		} else if Section == RODATA {
			_, err := rTemp.Write(b)
			if err != nil {
				return 0, err
			}
		}
		offset = uint(len(b))
	} else {
		offset = 4
	}
	return offset, nil
}

func EncodeData(line []string) ([]byte, error) {
	var size int
	var err error = nil
	var data int64
	isString := false
	isData := false
	nullTerm := false

	line[1] = strings.ReplaceAll(line[1], "\"", "")

	switch line[0] {
	case ".space":
		size, err = strconv.Atoi(line[1])
	case ".word":
		size = 4
		isData = true
	case ".byte":
		size = 1
		isData = true
	case ".half":
		size = 2
		isData = true
	case ".asciiz":
		size = len(line[1]) + 1
		isString = true
		nullTerm = true
	case ".ascii":
		size = len(line[1])
		isString = true
	}
	if err != nil {
		return nil, err
	}
	bytes := make([]byte, size)

	if isData {
		if strings.Contains(line[1], "0x") {
			line[1] = strings.ReplaceAll(line[1], "0x", "")
			var t uint64
			t, err = strconv.ParseUint(line[1], 16, size*8)
			data = int64(t)
		} else {
			data, err = strconv.ParseInt(line[1], 10, size*8)
		}
		if err != nil {
			return nil, err
		}
		switch size {
		case 1:
			bytes[0] = uint8(data)
		case 2:
			binary.NativeEndian.PutUint16(bytes, uint16(data))
		case 4:
			binary.NativeEndian.PutUint32(bytes, uint32(data))
		}
	} else if isString {
		if nullTerm {
			for i, b := range []byte(line[1]) {
				bytes[i] = b
			}
			bytes[size-1] = 0
		} else {
			bytes = []byte(line[1])
		}
	}
	return bytes, nil
}
