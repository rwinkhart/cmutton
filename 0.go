package main

// #include <string.h>
// #include <stdlib.h>
// #include "types.h"
import "C"
import "unsafe"

// getPascalString returns a pascal
// string struct for the input Go string.
func getPascalString(goString string) C.PascalString {
	goStringBytes := []byte(goString)
	goStringPtr := C.CBytes(goStringBytes)
	return C.PascalString{
		data: (*C.char)(goStringPtr),
		len:  C.int(len(goStringBytes)),
	}
}

// getPascalStringFromBytes returns a pascal
// string struct for the input Go byte slice.
func getPascalStringFromBytes(goBytes []byte) C.PascalString {
	goBytesPtr := C.CBytes(goBytes)
	return C.PascalString{
		data: (*C.char)(goBytesPtr),
		len:  C.int(len(goBytes)),
	}
}

// GetPascalStringFromCString is a helper meant to be used from C.
//
//export GetPascalStringFromCString
func GetPascalStringFromCString(cString *C.char) C.PascalString {
	return C.PascalString{
		data: cString,
		len:  C.int(C.strlen(cString)),
	}
}

// FreeArray frees the memory allocated by a C-allocated array.
//
//export FreeArray
func FreeArray(items *C.PascalString, count C.int) {
	linesSlice := (*[1 << 30]C.PascalString)(unsafe.Pointer(items))[:count:count]
	for i := 0; i < int(count); i++ {
		C.free(unsafe.Pointer(linesSlice[i].data))
	}
	C.free(unsafe.Pointer(items))
}

// safeStringDeref safely dereferences a string pointer, returning empty string if nil.
func safeStringDeref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// safeBoolDeref safely dereferences a bool pointer, returning false if nil.
func safeBoolDeref(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}
