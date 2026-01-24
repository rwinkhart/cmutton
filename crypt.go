package main

// #include <stdlib.h>
// #include "types.h"
import "C"
import (
	"unsafe"

	"github.com/rwinkhart/libmutton/crypt"
	"github.com/rwinkhart/rcw/wrappers"
)

// DecryptFileToSlice returns:
// r0: err
//
// r1: decLines (pointer to C-allocated array)
//
// r2: length decLines array
//
//export DecryptFileToSlice
func DecryptFileToSlice(realPath, rcwPassword C.PascalString) (*C.char, *C.PascalString, C.int) {
	lines, err := crypt.DecryptFileToSlice(C.GoStringN(realPath.data, realPath.len), []byte(C.GoStringN(rcwPassword.data, rcwPassword.len)))
	if err != nil {
		return C.CString(err.Error()), nil, 0
	}

	if len(lines) == 0 {
		return nil, nil, 0
	}

	// allocate C array for PascalStrings
	cLinesPtr := (*C.PascalString)(C.malloc(C.size_t(len(lines)) * C.size_t(C.sizeof_PascalString)))

	// populate the array with the decrypted lines
	cLinesSlice := (*[1 << 30]C.PascalString)(unsafe.Pointer(cLinesPtr))[:len(lines):len(lines)]
	for i, line := range lines {
		cLinesSlice[i] = getPascalString(line)
	}

	return nil, cLinesPtr, C.int(len(lines))
}

// EncryptBytes returns:
// encBytes
//
//export EncryptBytes
func EncryptBytes(decBytes, rcwPassword C.PascalString) C.PascalString {
	// use wrappers.Encrypt directly since C bindings do no support the RCWD daemon
	encBytes := wrappers.Encrypt([]byte(C.GoStringN(decBytes.data, decBytes.len)), []byte(C.GoStringN(rcwPassword.data, rcwPassword.len)))
	return getPascalStringFromBytes(encBytes)
}
