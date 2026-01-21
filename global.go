package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"github.com/rwinkhart/libmutton/global"
)

// DirInit returns:
// r0: oldDeviceID (special case: see libmutton docs)
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

func main() {}
