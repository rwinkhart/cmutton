package main

// #include <stdlib.h>
// #include "types.h"
import "C"
import (
	"github.com/rwinkhart/libmutton/global"
)

// DirInit returns:
// r0: err
//
// r1: oldDeviceID (nil w/-1 length if none present)
//
//export DirInit
func DirInit(preserveOldCfgDir bool) (*C.char, C.PascalString) {
	oldDeviceID, err := global.DirInit(preserveOldCfgDir)
	if err != nil {
		return C.CString(err.Error()), C.PascalString{data: nil, len: 0}
	}
	if oldDeviceID != nil {
		return nil, getPascalString(*oldDeviceID)
	}
	return nil, C.PascalString{data: nil, len: -1}
}

// GetCurrentDeviceID returns:
// r0: err
//
// r1: currentDeviceID (nil w/-1 length if none present)
//
//export GetCurrentDeviceID
func GetCurrentDeviceID() (*C.char, C.PascalString) {
	currentDeviceID, err := global.GetCurrentDeviceID()
	if err != nil {
		return C.CString(err.Error()), C.PascalString{data: nil, len: 0}
	}
	if currentDeviceID != nil {
		return nil, getPascalString(*currentDeviceID)
	}
	return nil, C.PascalString{data: nil, len: -1}
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

// GetVersion returns:
// versionNumber
//
//export GetVersion
func GetVersion() *C.char {
	return C.CString(global.LibmuttonVersion)
}

// GetSSHDirPath returns:
// sshDirPath
//
//export GetSSHDirPath
func GetSSHDirPath() C.PascalString {
	return getPascalString(global.SSHDir)
}
