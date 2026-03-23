(module
	(import "wasi_snapshot_preview1" "fd_write" (func $fd_write (param i32 i32 i32 i32) (result i32)))
	(memory 1)
	(export "memory" (memory 0))
	;; Two i32 values take up exactly 8 bytes.
	(data (i32.const 8) "Hi Luna!\n")
	(func $main (export "_start")
		;; i32.store takes two arguments, offset and value
		;; char[] message
		(i32.store (i32.const 0) (i32.const 8)) ;; *char[]
		(i32.store (i32.const 4) (i32.const 9)) ;; len(char[])
		(call $fd_write
			(i32.const 1) ;; File descriptor 1 for stdout, maybe 2 is for stderr?
			(i32.const 0) ;; *ptrs[]
			(i32.const 1) ;; len(ptrs[])
			(i32.const 20) ;; Pointer of where the written contents will go. 4 * 2 + 9 = 17, then pad to the nearest larger minimum multiple of 4.
		)
		drop
	)
)
