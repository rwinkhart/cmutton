//go:build !ios

package main

// #include <string.h>
// #include <stdlib.h>
import "C"
import "unsafe"

// securePtrOverwriteAndFree securely overwrites the memory at the input pointer and then frees it.
func securePtrOverwriteAndFree(input unsafe.Pointer, length C.size_t) {
	C.explicit_bzero(input, C.size_t(length))
	C.free(input)
}
