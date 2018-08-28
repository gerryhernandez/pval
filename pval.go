package pval

func Bool(val bool) *bool {
	valCopy := *&val
	return &valCopy
}

func Bools(vals ...bool) []*bool {
	var result []*bool
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}

func String(val string) *string {
	valCopy := *&val
	return &valCopy
}

func Strings(vals ...string) []*string {
	var result []*string
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}

func Int(val int) *int {
	valCopy := *&val
	return &valCopy
}

func Ints(vals ...int) []*int {
	var result []*int
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}

func Uint8(val uint8) *uint8 {
	valCopy := *&val
	return &valCopy
}

func Uint8s(vals ...uint8) []*uint8 {
	var result []*uint8
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}

func Uint16(val uint16) *uint16 {
	valCopy := *&val
	return &valCopy
}

func Uint16s(vals ...uint16) []*uint16 {
	var result []*uint16
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}

func Uint32(val uint32) *uint32 {
	valCopy := *&val
	return &valCopy
}

func Uint32s(vals ...uint32) []*uint32 {
	var result []*uint32
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}

func Uint64(val uint64) *uint64 {
	valCopy := *&val
	return &valCopy
}

func Uint64s(vals ...uint64) []*uint64 {
	var result []*uint64
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}

func Uintptr(val uintptr) *uintptr {
	valCopy := *&val
	return &valCopy
}

func Uintptrs(vals ...uintptr) []*uintptr {
	var result []*uintptr
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}

func Byte(val byte) *byte {
	valCopy := *&val
	return &valCopy
}

func Bytes(vals ...byte) []*byte {
	var result []*byte
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}

func Int8(val int8) *int8 {
	valCopy := *&val
	return &valCopy
}

func Int8s(vals ...int8) []*int8 {
	var result []*int8
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}

func Int16(val int16) *int16 {
	valCopy := *&val
	return &valCopy
}

func Int16s(vals ...int16) []*int16 {
	var result []*int16
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}

func Int32(val int32) *int32 {
	valCopy := *&val
	return &valCopy
}

func Int32s(vals ...int32) []*int32 {
	var result []*int32
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}

func Int64(val int64) *int64 {
	valCopy := *&val
	return &valCopy
}

func Int64s(vals ...int64) []*int64 {
	var result []*int64
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}

func Rune(val rune) *rune {
	valCopy := *&val
	return &valCopy
}

func Runes(vals ...rune) []*rune {
	var result []*rune
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}

func Float32(val float32) *float32 {
	valCopy := *&val
	return &valCopy
}

func Float32s(vals ...float32) []*float32 {
	var result []*float32
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}

func Float64(val float64) *float64 {
	valCopy := *&val
	return &valCopy
}

func Float64s(vals ...float64) []*float64 {
	var result []*float64
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}

func Complex64(val complex64) *complex64 {
	valCopy := *&val
	return &valCopy
}

func Complex64s(vals ...complex64) []*complex64 {
	var result []*complex64
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}

func Complex128(val complex128) *complex128 {
	valCopy := *&val
	return &valCopy
}

func Complex128s(vals ...complex128) []*complex128 {
	var result []*complex128
	for _, val := range vals {
		valCopy := *&val
		result = append(result, &valCopy)
	}
	return result
}
