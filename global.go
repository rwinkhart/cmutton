package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"github.com/rwinkhart/libmutton/global"
)

// DirInit returns:
// r0: oldDeviceID (FSMisc if none)
//
// r1: err
//
//export DirInit
func DirInit(preserveOldCfgDir bool) (*C.char, *C.char) {
	oldDeviceID, err := global.DirInit(preserveOldCfgDir)
	if err != nil {
		return nil, C.CString(err.Error())
	}
	return C.CString(oldDeviceID), nil
}

// GetCurrentDeviceID returns:
// r0: currentDeviceID (FSMisc if none)
//
// r1: err
//
//export GetCurrentDeviceID
func GetCurrentDeviceID() (*C.char, *C.char) {
	currentDeviceID, err := global.GetCurrentDeviceID()
	if err != nil {
		return nil, C.CString(err.Error())
	}
	return C.CString(currentDeviceID), nil
}

// GetRealAgePath returns:
// realAgePath
//
//export GetRealAgePath
func GetRealAgePath(vanityPath *C.char) *C.char {
	return C.CString(global.GetRealAgePath(C.GoString(vanityPath)))
}

// GetRealPath returns:
// realPath
//
//export GetRealPath
func GetRealPath(vanityPath *C.char) *C.char {
	return C.CString(global.GetRealPath(C.GoString(vanityPath)))
}

// GetVanityPath returns:
// vanityPath
//
//export GetVanityPath
func GetVanityPath(realPath *C.char) *C.char {
	return C.CString(global.GetVanityPath(C.GoString(realPath)))
}

func main() {}
