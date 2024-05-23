package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var isText bool = true

var Instructions = map[string]struct {
	isRtype bool
	opcode  uint32
}{
	"add":     {true, 32},
	"addi":    {false, 8},
	"addiu":   {false, 9},
	"sub":     {true, 34},
	"subu":    {true, 35},
	"mult":    {true, 24},
	"multu":   {true, 25},
	"div":     {true, 26},
	"divu":    {true, 27},
	"and":     {true, 36},
	"andi":    {false, 12},
	"or":      {true, 37},
	"ori":     {false, 13},
	"xor":     {true, 38},
	"xori":    {false, 14},
	"sll":     {true, 0},
	"sra":     {true, 3},
	"srl":     {true, 2},
	"lw":      {false, 35},
	"sw":      {false, 43},
	"slt":     {true, 42},
	"slti":    {false, 10},
	"sltiu":   {false, 11},
	"sltu":    {true, 43},
	"beq":     {false, 4},
	"bne":     {false, 5},
	"j":       {false, 2},
	"jal":     {false, 3},
	"jr":      {true, 8},
	"syscall": {false, 12},
}

var RegNums = map[string]uint32{
	"$zero": 0,
	"$at":   1,
	"$v0":   2,
	"$v1":   3,
	"$a0":   4,
	"$a1":   5,
	"$a2":   6,
	"$a3":   7,
	"$t0":   8,
	"$t1":   9,
	"$t2":   10,
	"$t3":   11,
	"$t4":   12,
	"$t5":   13,
	"$t6":   14,
	"$t7":   15,
	"$s0":   16,
	"$s1":   17,
	"$s2":   18,
	"$s3":   19,
	"$s4":   20,
	"$s5":   21,
	"$s6":   22,
	"$s7":   23,
	"$t8":   24,
	"$t9":   25,
	"$k0":   26,
	"$k1":   27,
	"$gp":   28,
	"$sp":   29,
	"$fp":   30,
	"$ra":   31,
}

func Encode(inst []string, pc int32) ([]byte, error) {
	var ret uint32 = 0
	var imm int16
	bytes := make([]byte, 4)
	for _, s := range inst {
		switch s {
		case "syscall":
			binary.NativeEndian.PutUint32(bytes, 12)
			return bytes, nil
		case "nop":
			binary.NativeEndian.PutUint32(bytes, 0)
			return bytes, nil
		}
	}

	function := Instructions[inst[0]]
	if function.isRtype && function.opcode == 8 {
		ret = (RegNums[inst[1]] << 21) | function.opcode
	} else if function.isRtype {
		ret = (RegNums[inst[2]] << 21) | (RegNums[inst[3]] <<
			16) | (RegNums[inst[1]] << 11) | function.opcode
	} else if function.opcode == 2 || function.opcode == 3 {
		label, _ := labels[inst[1]]
		ret = (function.opcode << 26) | uint32(label&0x03FFFFFF)
	} else if function.opcode == 35 || function.opcode == 43 {
		immReg, err := regexp.Compile("^[0-9]+")
		if err != nil {
			return nil, err
		}
		regReg, err := regexp.Compile("\\$([a-z]|[0-9])+")
		if err != nil {
			return nil, err
		}
		t, err := strconv.Atoi(immReg.FindString(inst[2]))
		if err != nil {
			return nil, err
		}
		imm = int16(t)
		reg := regReg.FindString(inst[2])
		ret = (function.opcode << 26) | (RegNums[reg] << 21) |
			(RegNums[inst[1]] << 16) | (0xFFFF & uint32(imm))
	} else {
		addr, isLabel := labels[inst[3]]
		if isLabel {
			imm = int16(int32(addr) - pc - 12)
		} else {
			i, err := strconv.Atoi(inst[3])
			if err != nil {
				return nil, err
			}
			imm = int16(i)
		}
		ret = (function.opcode << 26) | (RegNums[inst[2]] << 21) |
			(RegNums[inst[1]] << 16) | (0xFFFF & uint32(imm))
	}
	binary.NativeEndian.PutUint32(bytes, ret)
	return bytes, nil
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

func Assemble(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	out, err := os.Create(strings.ReplaceAll(path, ".asm", ".bin"))
	if err != nil {
		panic(err)
	}
	defer out.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i += 4 {
		var s string = scanner.Text()
		if s == "" {
			i -= 4
			continue
		} else if s == ".data" {
			isText = false
			continue
		} else if s == ".text" {
			isText = true
			continue
		}

		inst, err := Parse(s)
		if err != nil {
			panic(err)
		}

		_, isLabel := labels[strings.ReplaceAll(inst[0], ":", "")]
		if isLabel {
			inst = inst[1:]
		}

		if len(inst) == 0 || (len(inst) == 1 && inst[0] == "") {
			i -= 4
			continue
		}

		var bytes []byte
		if isText {
			bytes, err = Encode(inst, int32(i))
		} else {
			bytes, err = EncodeData(inst)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "while encoding %s: %v", s, err)
			panic(err)
		}
		if _, err = out.Write(bytes); err != nil {
			panic(err)
		}
	}

	/*
		// Write the new starting PC into the header
		if _, err := out.Seek(0, io.SeekStart); err != nil {
			panic(err)
		}
	*/

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}