# This is a test program to test the capabilities of the assembler

addi $t0, $zero, 10		# load 10 into $t0
addi $t1, $zero, 12		# load 12 into $t1
add $t2, $t1, $t0
jal PrintInt
nop
j Exit # Exit the program
nop

PrintInt: add $a0, $t2, $zero
addi $v0, $zero, 1
syscall
jr $ra
nop

Exit: addi $v0, $zero, 10
syscall
