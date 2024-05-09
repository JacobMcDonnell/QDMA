addi $t0, $zero, 10
addi $t1, $zero, 12
add $t2, $t1, $t0

add $a0, $t2, $zero
addi $v0, $zero, 1
syscall

addi $v0, $zero, 10
syscall
