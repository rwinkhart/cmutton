package main

// #include <stdlib.h>
// #include "types.h"
import "C"
import (
	"github.com/rwinkhart/libmutton/syncclient"
)

// AddFolderRemote returns:
// r0: err
//
//export AddFolderRemote
func AddFolderRemote(vanityPath C.PascalString) *C.char {
	if err := syncclient.AddFolderRemote(C.GoStringN(vanityPath.data, vanityPath.len)); err != nil {
		return C.CString(err.Error())
	}
	return nil
}

// RenameRemote returns:
// r0: err
//
//export RenameRemote
func RenameRemote(oldVanityPath, newVanityPath C.PascalString) *C.char {
	if err := syncclient.RenameRemote(C.GoStringN(oldVanityPath.data, oldVanityPath.len), C.GoStringN(newVanityPath.data, newVanityPath.len)); err != nil {
		return C.CString(err.Error())
	}
	return nil
}

// ShearRemote returns:
// r0: err
//
//export ShearRemote
func ShearRemote(vanityPath C.PascalString, onlyShearAgeFile bool) *C.char {
	if err := syncclient.ShearRemote(C.GoStringN(vanityPath.data, vanityPath.len), onlyShearAgeFile); err != nil {
		return C.CString(err.Error())
	}
	return nil
}

// GenDeviceID returns:
// r0: err
//
// r1: remoteEntryRoot
//
// r2: remoteAgeDir
//
// r3: isWindows
//
//export GenDeviceID
func GenDeviceID(oldDeviceID, prefix C.PascalString) (*C.char, C.PascalString, C.PascalString, bool) {
	var remoteEntryRoot, remoteAgeDir string
	var isWindows bool
	var err error
	if oldDeviceID.len < 1 {
		remoteEntryRoot, remoteAgeDir, isWindows, err = syncclient.GenDeviceID(nil, C.GoStringN(prefix.data, prefix.len))
	} else {
		remoteEntryRoot, remoteAgeDir, isWindows, err = syncclient.GenDeviceID(new(C.GoStringN(oldDeviceID.data, oldDeviceID.len)), C.GoStringN(prefix.data, prefix.len))
	}
	if err != nil {
		return C.CString(err.Error()), C.PascalString{data: nil, len: 0}, C.PascalString{data: nil, len: 0}, false
	}
	return nil, getPascalString(remoteEntryRoot), getPascalString(remoteAgeDir), isWindows
}

// RunJob returns:
// r0: err
//
// r1: deleteList (pointer to C-allocated array)
//
// r2: deleteList length
//
// r3: downloadList (pointer to C-allocated array)
//
// r4: downloadList length
//
// r5: uploadList (pointer to C-allocated array)
//
// r6: uploadList length
//
//export RunJob
func RunJob() (*C.char, *C.PascalString, C.int, *C.PascalString, C.int, *C.PascalString, C.int) {
	lists, err := syncclient.RunJob()
	if err != nil {
		return C.CString(err.Error()), nil, 0, nil, 0, nil, 0
	}
	return nil, getCPascalStringArrayFromStringSlice(lists.Delete), C.int(len(lists.Delete)), getCPascalStringArrayFromStringSlice(lists.Download), C.int(len(lists.Download)), getCPascalStringArrayFromStringSlice(lists.Upload), C.int(len(lists.Upload))
}
