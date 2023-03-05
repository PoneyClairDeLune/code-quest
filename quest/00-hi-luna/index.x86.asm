# Template from Low Level Learning

.global _start
.intel_syntax
.section .text

_start:
	
	# Print
	mov %eax, 4
	mov %ebx, 1
	lea %ecx, [message]
	mov %edx, 9
	int 0x80
	
	# Exit
	mov %eax, 1
	mov %ebx, 0
	int 0x80

.section .data
	message:
	.ascii "Hi Luna!\n"
