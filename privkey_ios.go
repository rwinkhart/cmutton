//go:build ios

package main

// #include <stdlib.h>
// #include <string.h>
// #include "types.h"
import "C"
import (
	"unsafe"

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
	privkey.LocalBytes = C.GoBytes(privKeyDataPtr, privKeyData.len)
	securePtrOverwriteAndFree(privKeyDataPtr, C.size_t(privKeyData.len))
	FreePascalString(privKeyData)
}
