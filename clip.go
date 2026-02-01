//go:build (!android && !ios) || termux

package main

// #include <stdlib.h>
// #include "types.h"
import "C"
import (
	"github.com/rwinkhart/libmutton/clip"
)

// ClearProcess clears the clipboard in 30 seconds if its contents still match assignedContents.
// If assignedContents is an empty string, the clipboard will clear immediately and unconditionally.
// Returns:
// r0: err
//
//export ClearProcess
func ClearProcess(assignedContents C.PascalString) *C.char {
	if err := clip.ClearProcess(C.GoStringN(assignedContents.data, assignedContents.len)); err != nil {
		return C.CString(err.Error())
	}
	return nil
}

// CopyShortcut returns:
// r0: err
//
//export CopyShortcut
func CopyShortcut(realPath C.PascalString, field int, rcwPassword C.PascalString) *C.char {
	if err := clip.CopyShortcut(C.GoStringN(realPath.data, realPath.len), field, []byte(C.GoStringN(rcwPassword.data, rcwPassword.len))); err != nil {
		return C.CString(err.Error())
	}
	return nil
}

// CopyString returns:
// r0: err
//
//export CopyString
func CopyString(clearClipboardAutomatically bool, copySubject C.PascalString) *C.char {
	if err := clip.CopyString(clearClipboardAutomatically, C.GoStringN(copySubject.data, copySubject.len)); err != nil {
		return C.CString(err.Error())
	}
	return nil
}
