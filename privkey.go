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

// GetPrivKeyDataAndEraseGoCopy will return nil on non-iOS platforms.
// On iOS, it will return whatever was set in
// SetPrivKeyDataAndSecurelyFreeInput.
//
// Don't forget to call EraseFreePrivKeyPascal on the output
// of this function as soon as it is no longer needed.
//
//export GetPrivKeyDataAndEraseGoCopy
func GetPrivKeyDataAndEraseGoCopy() C.PascalString {
	privKeyBytes, _ := privkey.GetBytes(nil)
	privKeyPascal := getPascalStringFromBytes(privKeyBytes)
	back.EraseBytesSecurely(privKeyBytes)
	return privKeyPascal
}

// EraseFreePrivKeyPascal is meant to be called on
// iOS after any server connections to ensure the cached SSH private key
// set by SetPrivKeyData is securely erased. It also erases and frees
// the input privKeyData C.PascalString.
//
//export EraseFreePrivKeyPascal
func EraseFreePrivKeyPascal(privKeyPascal C.PascalString) {
	securePtrOverwriteAndFree(unsafe.Pointer(privKeyPascal.data), C.size_t(privKeyPascal.len))
	FreePascalString(privKeyPascal)
}
