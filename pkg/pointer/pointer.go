package pointer

// IntPtr returns *int
func IntPtr(v int) *int {
	return &v
}

// Int64Ptr returns *int64
func Int64Ptr(v int64) *int64 {
	return &v
}

// UintPtr returns *uint
func UintPtr(v uint) *uint {
	return &v
}

// Uint64Ptr returns *uint64
func Uint64Ptr(v uint64) *uint64 {
	return &v
}

// Float64Ptr returns *float64
func Float64Ptr(v float64) *float64 {
	return &v
}

// BoolPtr returns *bool
func BoolPtr(v bool) *bool {
	return &v
}

// StringPtr returns *string
func StringPtr(v string) *string {
	return &v
}

// StringValues returns values of the *string list
func StringValues(ptrs []*string) []string {
	values := make([]string, len(ptrs))
	for i := 0; i < len(ptrs); i++ {
		if ptrs[i] != nil {
			values[i] = *ptrs[i]
		}
	}
	return values
}

// IntPtrs returns *int list
func IntPtrs(vals []int) []*int {
	ptrs := make([]*int, len(vals))
	for i := 0; i < len(vals); i++ {
		ptrs[i] = &vals[i]
	}
	return ptrs
}

// Int64Ptrs returns *int64 list
func Int64Ptrs(vals []int64) []*int64 {
	ptrs := make([]*int64, len(vals))
	for i := 0; i < len(vals); i++ {
		ptrs[i] = &vals[i]
	}
	return ptrs
}

// UintPtrs returns *uint list
func UintPtrs(vals []uint) []*uint {
	ptrs := make([]*uint, len(vals))
	for i := 0; i < len(vals); i++ {
		ptrs[i] = &vals[i]
	}
	return ptrs
}

// Uint64Ptrs returns *uint64 list
func Uint64Ptrs(vals []uint64) []*uint64 {
	ptrs := make([]*uint64, len(vals))
	for i := 0; i < len(vals); i++ {
		ptrs[i] = &vals[i]
	}
	return ptrs
}

// Float64Ptrs returns *float64 list
func Float64Ptrs(vals []float64) []*float64 {
	ptrs := make([]*float64, len(vals))
	for i := 0; i < len(vals); i++ {
		ptrs[i] = &vals[i]
	}
	return ptrs
}

// BoolPtrs returns *bool list
func BoolPtrs(vals []bool) []*bool {
	ptrs := make([]*bool, len(vals))
	for i := 0; i < len(vals); i++ {
		ptrs[i] = &vals[i]
	}
	return ptrs
}

// StringPtrs returns *string list
func StringPtrs(vals []string) []*string {
	ptrs := make([]*string, len(vals))
	for i := 0; i < len(vals); i++ {
		ptrs[i] = &vals[i]
	}
	return ptrs
}
