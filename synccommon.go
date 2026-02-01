package main

// #include <stdlib.h>
// #include "types.h"
import "C"
import (
	"unsafe"

	"github.com/rwinkhart/libmutton/age"
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

// GetAllEntryData returns:
// r0: err
//
// r1: vanityPaths (pointer to C-allocated array)
//
// r2: vanityPaths length
//
// r3: translatedAges (pointer to C-allocated array)
//
// r4: translatedAges length
//
//export GetAllEntryData
func GetAllEntryData() (*C.char, *C.PascalString, C.int, *uint8, C.int) {
	entryMap, err := synccommon.GetAllEntryData()
	if err != nil {
		return C.CString(err.Error()), nil, 0, nil, 0
	}
	var vanityPaths []string
	var translatedAges []uint8
	for k, v := range entryMap {
		vanityPaths = append(vanityPaths, k)
		translatedAges = append(translatedAges, age.TranslateAgeTimestamp(v.AgeTimestamp))
	}
	return nil, getCPascalStringArrayFromStringSlice(vanityPaths), C.int(len(vanityPaths)), (*uint8)(unsafe.Pointer(&translatedAges[0])), C.int(len(translatedAges))
}
