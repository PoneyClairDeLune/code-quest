(module
	(func (export "readBitAt") (param i32 i32) (result i32)
		;; The first parameter should be the incoming bit-field.
		;; The second parameter should be the position where the bit lives.
		local.get 0
		local.get 1
		i32.const 7
		i32.and
		i32.shr_u
		i32.const 1
		i32.and
	)
)