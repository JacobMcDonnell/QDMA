.text

addi $t0, $zero, x
addi $v0, $zero, 1
lw $a0, 0($t0)
syscall
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

