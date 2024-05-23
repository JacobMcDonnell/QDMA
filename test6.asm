.text

addi $t0, $zero, x
addi $v0, $zero, 1
lw $a0, 0($t0)
syscall
addi $v0, $zero, 10
syscall

.data

x: .word 0xFFFFFFFF
#fib: .space 40
#h: .half 255
#c: .byte 0
#dogs: .ascii "Dogs are cool\n"
#cats: .asciiz "cats are cool\n"

