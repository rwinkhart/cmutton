package main

// #include <stdlib.h>
// #include <string.h>
// #include "types.h"
import "C"
import (
	"unsafe"

	"github.com/rwinkhart/go-boilerplate/back"
	"github.com/rwinkhart/libmutton/privkey"
)

// SetPrivKeyDataAndEraseFreeInput is meant to be called on iOS
// prior to any server connections to ensure libmutton can
// retrieve the correct SSH private key contents (enables
// loading from keychain).
//
//export SetPrivKeyDataAndEraseFreeInput
func SetPrivKeyDataAndEraseFreeInput(privKeyData C.PascalString) {
	privKeyDataPtr := unsafe.Pointer(privKeyData.data)
	privkey.SetBytes(C.GoBytes(privKeyDataPtr, privKeyData.len))
	securePtrOverwriteAndFree(privKeyDataPtr, C.size_t(privKeyData.len))
	FreePascalString(privKeyData)
}

// EraseFreePrivKey is meant to be called on iOS after
// any server connections to ensure the cached SSH private key
// set by SetPrivKeyDataAndEraseFreeInput is securely erased.
// It also erases and frees the input privKeyPascal.
//
//export EraseFreePrivKey
func EraseFreePrivKey(privKeyPascal C.PascalString) {
	// erase Go/libmutton copy
	privKeyBytes, _ := privkey.GetBytes(nil)
	back.EraseBytesSecurely(privKeyBytes)

	// erase C copy
	securePtrOverwriteAndFree(unsafe.Pointer(privKeyPascal.data), C.size_t(privKeyPascal.len))
	FreePascalString(privKeyPascal)
}
