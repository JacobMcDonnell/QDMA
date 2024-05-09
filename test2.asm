.data

fib: space 40

str: .asciiz "Hello"

.text
main:	addi $t0, $zero, 10
	addi $t1, $zero, 11
	loop:	add $t2, $t1, $t0
	
		addi $v0, $zero, 1
		addi $a0, $t2, 0
		syscall

		addi $t0, $t0, -1
		addi $t1, $t1, -1
		bne $t0, $zero, loop
	addi $v0, $zero, 10
	syscall

