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

func main() {}
