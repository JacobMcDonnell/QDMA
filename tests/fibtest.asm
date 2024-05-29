.data

str:
	.asciiz "Fibinacci Number "
	
str2:
	.asciiz " is: "

fib:
	.space 40
	
.text

main:
	addi $t0, $zero, 0
	addi $t1, $zero, 1
	add $t3, $zero, $zero
	la $t6, fib
	addi $t5, $zero, 9
	
	loop:
		add $t2, $t1, $t0
		add $t1, $zero, $t0
		add $t0, $zero, $t2
		sw $t2, 00($t6)
		addi $t3, $t3, 1
		addi $t6, $t6, 4
		bne $t3, $t5, loop
		nop
	addi $t6, $t6, -4
	loop2:
		la $a0, str
		jal PrintStr
		nop
		
		add $a0, $zero, $t3
		jal PrintInt
		nop
		
		la $a0, str2
		jal PrintStr
		nop
		
		lw $a0, 00($t6)
		jal PrintInt
		nop
		
		addi $a0, $zero, 10
		addi $v0, $zero, 11
		syscall
		
		addi $t3, $t3, -1
		addi $t6, $t6, -4
		bne $t3, $zero, loop2
		nop
	
	addi $v0, $zero, 10
	syscall
	
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