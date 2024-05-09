package main

import (
	"bufio"
	"encoding/binary"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func Encode(instruction string) uint32 {
	var ret uint32 = 0
	if instruction == "syscall" {
		return 12
	}
	inst := strings.Fields(strings.ReplaceAll(instruction, ",", ""))
	if Instructions[inst[0]].isRtype {
		ret = (RegNums[inst[2]] << 21) | (RegNums[inst[3]] << 16) |
			(RegNums[inst[1]] << 11) | Instructions[inst[0]].opcode
	} else {
		var imm int32
		i, err := strconv.Atoi(inst[3])
		if err != nil {
			log.Fatal(err)
		}
		imm = int32(i)
		ret = (Instructions[inst[0]].opcode << 26) | (RegNums[inst[2]] << 21) |
			(RegNums[inst[1]] << 16) | uint32(imm)
	}
	return ret
}

func Assemble(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	out, err := os.Create(strings.ReplaceAll(path, ".asm", ".bin"))
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var s string = scanner.Text()
		if s != "" {
			bs := make([]byte, 4)
			binary.BigEndian.PutUint32(bs, Encode(s))
			_, err := out.Write(bs)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	for _, arg := range os.Args[1:] {
		Assemble(arg)
	}
}
