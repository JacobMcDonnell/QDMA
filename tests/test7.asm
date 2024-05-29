.text

la $t0, x
lw $a0, 0($t0)
jal PrintInt
nop
la $a0, cats
jal PrintStr
nop
j exit
nop

PrintInt:
	addi $v0, $zero, 1
	syscall
	jr $ra
	nop

PrintStr:
	addi $v0, $zero, 4
	syscall
	jr $ra
	nop

exit:
	addi $v0, $zero, 10
	syscall

.data

x: .word 0xFFFFFFFF
cats: .asciiz "cats are cool\n"
c: .byte 0

.rodata

g: .word 50
dogs: .ascii "Dogs are cool\n"
h: .half 255

.bss
fib: .space 40

