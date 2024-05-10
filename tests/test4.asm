addi $t0, $zero, 10
addi $t1, $zero, 12
add $t2, $t1, $t0
beq $zero, $zero, PrintInt
nop
caller: beq $zero, $zero, Exit
nop

PrintInt: add $a0, $t2, $zero
addi $v0, $zero, 1
syscall
beq $zero, $zero, caller
nop

Exit: addi $v0, $zero, 10
syscall
