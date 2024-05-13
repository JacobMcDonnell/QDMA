main: addiu   $sp,$sp,-40
sw      $31,36($sp)
sw      $fp,32($sp)
sw      $16,28($sp)
move    $fp,$sp
sw      $4,40($fp)
j main
