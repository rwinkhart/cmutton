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

// getCPascalStringArrayFromStringSlice converts a Go slice
// of strings to a C array of pascal strings.
func getCPascalStringArrayFromStringSlice(goSlice []string) C.PascalStringArray {
	// allocate C array for PascalStrings
	cArrPtr := (*C.PascalString)(C.malloc(C.size_t(len(goSlice)) * C.size_t(C.sizeof_PascalString)))

	// populate the array with the decrypted lines
	cArr := (*[1 << 30]C.PascalString)(unsafe.Pointer(cArrPtr))[:len(goSlice):len(goSlice)]
	for i := range goSlice {
		cArr[i] = getPascalString(goSlice[i])
	}

	return C.PascalStringArray{data: cArrPtr, len: C.int(len(goSlice))}
}

// getStringSliceFromCPascalStringArray converts a C.PascalStringArray to a go slice.
func getStringSliceFromCPascalStringArray(pascalStringArray C.PascalStringArray) []string {
	var goSlice []string
	cArr := (*[1 << 30]C.PascalString)(unsafe.Pointer(pascalStringArray.data))[:pascalStringArray.len:pascalStringArray.len]
	for i := 0; i < int(pascalStringArray.len); i++ {
		goSlice = append(goSlice, C.GoStringN(cArr[i].data, cArr[i].len))
	}
	return goSlice
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
func FreeArray(array C.PascalStringArray) {
	linesSlice := (*[1 << 30]C.PascalString)(unsafe.Pointer(array.data))[:array.len:array.len]
	for i := 0; i < int(array.len); i++ {
		C.free(unsafe.Pointer(linesSlice[i].data))
	}
	C.free(unsafe.Pointer(array.data))
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
