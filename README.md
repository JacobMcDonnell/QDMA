# QDMA - Quick and Dirty MIPS Assembler

## Intro

QDMA is a simple MIPS assembler written in Go for
[QDME](https://github.com/JacobMcDonnell/QDME), which is my simple MIPS
emulator. QDMA implements a partial version of the Executable and Linkable
Format (ELF), [See Partial ELF for more information](#Partial-ELF).

## TODO

- ~Assemble Instructions into Binary~

- ~Handle Labels and Jumps~

- ~Encode Data Section into Binary~

- ~Ignore Comments~

## Instructions

Here is a list of supported instructions.

|Instruction|Type|Action|
|-----------|----|------|
|add        |R   |rd = rs + rt|
|addu       |R   |rd = rs + rt|
|and        |R   |rd = rs & rt|
|or         |R   |rd = rs \| rt|
|slt        |R   |rd = rs < rt|
|sltu       |R   |rd = rs < rt|
|sub        |R   |rd = rs - rt|
|subu       |R   |rd = rs - rt|
|xor        |R   |rd = rs ^ rt|
|sll        |R   |rd = rt << shamt|
|sra        |R   |rd = rt >> shamt|
|srl        |R   |rd = rt >> shamt|
|div        |R   |hi = rs % rt lo = rs / rt|
|divu       |R   |hi = rs % rt lo = rs / rt|
|mult       |R   |hi, lo = rs * rt|
|multu      |R   |hi, lo = rs * rt|
|MFHI       |R   |rd = hi|
|MFLO       |R   |rd = lo|
|MTHI       |R   |rs = hi|
|MTLO       |R   |rs = lo|
|jr         |R   |pc = addr, pc = $ra|
|jalr       |R   |rd = pc, pc = addr|
|syscall    |R   |System Call|
|jal        |J   |$ra = pc, pc = addr|
|j          |J   |pc = addr|
|beq        |I   | pc = (rs == rt) ? imm + pc + 4 : pc|
|bne        |I   | pc = (rs != rt) ? imm + pc + 4 : pc|
|addi       |I   |rt = rs + imm|
|addiu      |I   |rt = rs + imm|
|andi       |I   |rt = rs & imm|
|ori        |I   |rt = rs \| imm|
|xori       |I   |rt = rs ^ imm|
|slti       |R   |rt = rs < imm|
|sltiu      |R   |rt = rs < imm|
|lui        |I   |rt = rs << 16|
|lw         |I   |rs = imm(rt)|
|sw         |I   |imm(rt) = rs|

## Partial ELF

QDMA is partiallialy complient with ELF, meaning that it supports segments but
not sections. The supported segments are `.text`, `.data`, `.rodata`, and
`.bss`. `.text` is for the instructions. `.data` and `.rodata` are for defined
data types with initial values, `.rodata` being read only. Supported data types
are `.word` for a 4 byte value, `.half` for a 2 byte value, `.byte` for a byte,
`.asciiz` for a null terminated string, and `.ascii` for a non null terminated
string. `.bss` is for variables where there is no initial value. In this
section you define the space that a variable will take up in bytes using
`.space`.
