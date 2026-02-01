package main

// #include <stdlib.h>
// #include "types.h"
import "C"
import (
	"github.com/rwinkhart/libmutton/synccommon"
)

// WalkEntryDir returns:
// r0: err
//
// r1: files (pointer to C-allocated array)
//
// r2: files length
//
// r3: dirs (pointer to C-allocated array)
//
// r4: dirs length
//
//export WalkEntryDir
func WalkEntryDir() (*C.char, *C.PascalString, C.int, *C.PascalString, C.int) {
	files, dirs, err := synccommon.WalkEntryDir()
	if err != nil {
		return C.CString(err.Error()), nil, 0, nil, 0
	}
	return nil, getCPascalStringArrayFromStringSlice(files), C.int(len(files)), getCPascalStringArrayFromStringSlice(dirs), C.int(len(dirs))
}
