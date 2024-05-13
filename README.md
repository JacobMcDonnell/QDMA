# QDMA - Quick and Dirty MIPS Assembler

## Intro

QDMA is a simple MIPS assembler written in Go for
[QDME](https://github.com/JacobMcDonnell/QDME), which is my simple MIPS
emulator. Currently this assembler can only handle standard instructions,
which means no sections and no predefined data outside of an I-Type
instruction.

## TODO

- ~Assemble Instructions into Binary~

- ~Handle Labels and Jumps~

- Encode Data Section into Binary

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

