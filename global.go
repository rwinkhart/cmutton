package main

// #include <stdlib.h>
// #include "types.h"
import "C"
import (
	"github.com/rwinkhart/libmutton/global"
)

// TODO Use a nil pointer for blank deviceIDs, not FSMisc (requires changes to libmutton)

// DirInit returns:
// r0: err
//
// r1: oldDeviceID (FSMisc if none)
//
//export DirInit
func DirInit(preserveOldCfgDir bool) (*C.char, C.PascalString) {
	oldDeviceID, err := global.DirInit(preserveOldCfgDir)
	if err != nil {
		return C.CString(err.Error()), C.PascalString{data: nil, len: 0}
	}
	return nil, getPascalString(oldDeviceID)
}

// GetCurrentDeviceID returns:
// r0: err
//
// r1: currentDeviceID (FSMisc if none)
//
//export GetCurrentDeviceID
func GetCurrentDeviceID() (*C.char, C.PascalString) {
	currentDeviceID, err := global.GetCurrentDeviceID()
	if err != nil {
		return C.CString(err.Error()), C.PascalString{data: nil, len: 0}
	}
	return nil, getPascalString(currentDeviceID)
}

// GetRealAgePath returns:
// realAgePath
//
//export GetRealAgePath
func GetRealAgePath(vanityPath *C.char) C.PascalString {
	return getPascalString(global.GetRealAgePath(C.GoString(vanityPath)))
}

// GetRealPath returns:
// realPath
//
//export GetRealPath
func GetRealPath(vanityPath *C.char) C.PascalString {
	return getPascalString(global.GetRealPath(C.GoString(vanityPath)))
}

// GetVanityPath returns:
// vanityPath
//
//export GetVanityPath
func GetVanityPath(realPath *C.char) C.PascalString {
	return getPascalString(global.GetVanityPath(C.GoString(realPath)))
}

func main() {}
