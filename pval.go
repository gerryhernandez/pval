package pval

func String(val string) *string {
	return &val
}

func Strings(vals ...string) []*string {
	var result []*string
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}

func Int(val int) *int {
	return &val
}

func Ints(vals ...int) []*int {
	var result []*int
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}

func Uint8(val uint8) *uint8 {
	return &val
}

func Uint8s(vals ...uint8) []*uint8 {
	var result []*uint8
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}

func Uint16(val uint16) *uint16 {
	return &val
}

func Uint16s(vals ...uint16) []*uint16 {
	var result []*uint16
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}

func Uint32(val uint32) *uint32 {
	return &val
}

func Uint32s(vals ...uint32) []*uint32 {
	var result []*uint32
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}

func Uint64(val uint64) *uint64 {
	return &val
}

func Uint64s(vals ...uint64) []*uint64 {
	var result []*uint64
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}

func Uintptr(val uintptr) *uintptr {
	return &val
}

func Uintptrs(vals ...uintptr) []*uintptr {
	var result []*uintptr
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}

func Byte(val byte) *byte {
	return &val
}

func Bytes(vals ...byte) []*byte {
	var result []*byte
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}

func Int8(val int8) *int8 {
	return &val
}

func Int8s(vals ...int8) []*int8 {
	var result []*int8
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}

func Int16(val int16) *int16 {
	return &val
}

func Int16s(vals ...int16) []*int16 {
	var result []*int16
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}

func Int32(val int32) *int32 {
	return &val
}

func Int32s(vals ...int32) []*int32 {
	var result []*int32
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}

func Int64(val int64) *int64 {
	return &val
}

func Int64s(vals ...int64) []*int64 {
	var result []*int64
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}

func Rune(val rune) *rune {
	return &val
}

func Runes(vals ...rune) []*rune {
	var result []*rune
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}

func Float32(val float32) *float32 {
	return &val
}

func Float32s(vals ...float32) []*float32 {
	var result []*float32
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}

func Float64(val float64) *float64 {
	return &val
}

func Float64s(vals ...float64) []*float64 {
	var result []*float64
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}

func Complex64(val complex64) *complex64 {
	return &val
}

func Complex64s(vals ...complex64) []*complex64 {
	var result []*complex64
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}

func Complex128(val complex128) *complex128 {
	return &val
}

func Complex128s(vals ...complex128) []*complex128 {
	var result []*complex128
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}
