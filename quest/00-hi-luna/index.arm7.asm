.global _start
.section .text

_start:
	# Print
	# "write"
	mov r7, #0x4
	# to stdout
	mov r0, #1
	# message address
	ldr r1, =message
	# 9 characters
	mov r2, #9
	# interrupt
	swi 0

	# Exit
	# "exit"
	mov r7, #0x1
	# returns 0
	mov r0, #0
	# interrupt
	swi 0
.section .data
	message:
	.ascii "Hi Luna!\n"
