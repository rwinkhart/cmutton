//go:build (!android && !ios) || termux

package main

// #include <stdlib.h>
// #include "types.h"
import "C"
import (
	"unsafe"

	"github.com/rwinkhart/libmutton/clip"
)

// ClearProcess clears the clipboard in 30 seconds if its contents still match assignedContents.
// If assignedContents is an empty string, the clipboard will clear immediately and unconditionally.
// Returns:
// r0: err
//
//export ClearProcess
func ClearProcess(assignedContents C.PascalString) *C.char {
	if err := clip.ClearProcess(C.GoBytes(unsafe.Pointer(assignedContents.data), assignedContents.len)); err != nil {
		return C.CString(err.Error())
	}
	return nil
}

// CopyShortcut returns:
// r0: err
//
//export CopyShortcut
func CopyShortcut(realPath C.PascalString, field int, rcwPassword C.PascalString) *C.char {
	if err := clip.CopyShortcut(C.GoStringN(realPath.data, realPath.len), field, C.GoBytes(unsafe.Pointer(rcwPassword.data), rcwPassword.len)); err != nil {
		return C.CString(err.Error())
	}
	return nil
}

// CopyString returns:
// r0: err
//
//export CopyString
func CopyString(clearClipboardAutomatically bool, copySubject C.PascalString) *C.char {
	if err := clip.CopyBytes(clearClipboardAutomatically, C.GoBytes(unsafe.Pointer(copySubject.data), copySubject.len)); err != nil {
		return C.CString(err.Error())
	}
	return nil
}
