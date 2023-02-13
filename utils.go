package spirv

import "unsafe"

// ptrToSlice casts an unsafe pointer & length to a go slice of a given type
// todo: replace with standard Go 1.20 implementation
func ptrToSlice[T any](ptr unsafe.Pointer, len int, cap int) []T {
	var sl = struct {
		addr unsafe.Pointer
		len  int
		cap  int
	}{ptr, len, cap}
	return *(*[]T)(unsafe.Pointer(&sl))
}
