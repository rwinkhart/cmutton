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
// r1: decLines
//
//export DecryptFileToSlice
func DecryptFileToSlice(realPath, rcwPassword C.PascalString) (*C.char, C.PascalStringArray) {
	lines, err := crypt.DecryptFileToSlice(C.GoStringN(realPath.data, realPath.len), C.GoBytes(unsafe.Pointer(rcwPassword.data), rcwPassword.len))
	if err != nil {
		return C.CString(err.Error()), C.PascalStringArray{}
	}

	if len(lines) == 0 {
		return nil, C.PascalStringArray{}
	}

	return nil, getCPascalStringArrayFromStringSlice(lines)
}

// EncryptBytes returns:
// encBytes
//
//export EncryptBytes
func EncryptBytes(decBytes, rcwPassword C.PascalString, zeroizeDecBytes, zeroizePassword bool) C.PascalString {
	// use wrappers.Encrypt directly since C bindings do not support the RCWD daemon
	encBytes := wrappers.Encrypt(C.GoBytes(unsafe.Pointer(decBytes.data), decBytes.len), C.GoBytes(unsafe.Pointer(rcwPassword.data), rcwPassword.len), zeroizeDecBytes, zeroizePassword)
	return getPascalStringFromBytes(encBytes)
}
