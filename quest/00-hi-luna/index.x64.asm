.global _start
.intel_syntax noprefix

_start:
	
	# Print
	# "write"
	mov rax, 1
	# to stdout
	mov rdi, 1
	# Load address of the message into memory
	lea rsi, [message]
	# Length of message: 9 characters
	mov rdx, 9
	# Syscall
	syscall
	
	# Exit
	# "exit"
	mov rax, 60
	# Returns 0
	mov rdi, 0
	# Syscall
	syscall

message:
	.asciz "Hi Luna!\n"
