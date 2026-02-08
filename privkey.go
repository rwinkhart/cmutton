package main

// #include <stdlib.h>
// #include "types.h"
import "C"
import (
	"unsafe"

	"github.com/rwinkhart/libmutton/privkey"
)

// SetPrivKeyBytes forces libmutton to use custom bytes for the SSH private key.
//
//export SetPrivKeyBytes
func SetPrivKeyBytes(privKeyPascal C.PascalString) {
	privkey.SetBytes(C.GoBytes(unsafe.Pointer(privKeyPascal.data), privKeyPascal.len))
}
