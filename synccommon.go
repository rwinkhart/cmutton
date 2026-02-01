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
// r1: files
//
// r2: dirs
//
//export WalkEntryDir
func WalkEntryDir() (*C.char, C.PascalStringArray, C.PascalStringArray) {
	files, dirs, err := synccommon.WalkEntryDir()
	if err != nil {
		return C.CString(err.Error()), C.PascalStringArray{}, C.PascalStringArray{}
	}
	return nil, getCPascalStringArrayFromStringSlice(files), getCPascalStringArrayFromStringSlice(dirs)
}

// GetAllEntryData returns:
// r0: err
//
// r1: vanityPaths
//
// r2: translatedAges (pointer to C-allocated array)
//
// r3: translatedAges length
//
//export GetAllEntryData
func GetAllEntryData() (*C.char, C.PascalStringArray, *uint8, C.int) {
	entryMap, err := synccommon.GetAllEntryData()
	if err != nil {
		return C.CString(err.Error()), C.PascalStringArray{}, nil, 0
	}
	var vanityPaths []string
	var translatedAges []uint8
	for k, v := range entryMap {
		vanityPaths = append(vanityPaths, k)
		translatedAges = append(translatedAges, age.TranslateAgeTimestamp(v.AgeTimestamp))
	}
	return nil, getCPascalStringArrayFromStringSlice(vanityPaths), (*uint8)(unsafe.Pointer(&translatedAges[0])), C.int(len(translatedAges))
}
