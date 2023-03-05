.global _start
.intel_syntax
.section .text

_start:
	
	# Print
	# "write"
	mov %eax, 1
	# to stdout
	mov %ebx, 1
	# Load address of the message into memory
	lea %ecx, [message]
	# Length of message: 9 characters
	mov %edx, 9
	# Interrupt
	int 0x80
	
	# Exit
	# "exit"
	mov %eax, 60
	# returns 0
	mov %ebx, 0
	# Interrupt
	int 0x80

.section .data
	message:
	.ascii "Hi Luna!\n"
