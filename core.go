package main

// #include <stdlib.h>
// #include "types.h"
import "C"
import (
	"strings"
	"time"

	"github.com/rwinkhart/libmutton/core"
)

// EntryAddPrecheck returns:
// r0: err
//
// r1: status
//
//export EntryAddPrecheck
func EntryAddPrecheck(realPath C.PascalString) (*C.char, uint8) {
	status, err := core.EntryAddPrecheck(C.GoStringN(realPath.data, realPath.len))
	if err != nil {
		return C.CString(err.Error()), status
	}
	return nil, status
}

// EntryIsNotEmpty returns:
// r0: notEmpty
//
//export EntryIsNotEmpty
func EntryIsNotEmpty(entryData C.PascalStringArray) bool {
	notEmpty := core.EntryIsNotEmpty(getStringSliceFromCPascalStringArray(entryData))
	return notEmpty
}

// EntryRefresh returns:
// r0: err
//
//export EntryRefresh
func EntryRefresh(oldRCWPassword, newRCWPassword C.PascalString, removeOldDir bool) *C.char {
	err := core.EntryRefresh([]byte(C.GoStringN(oldRCWPassword.data, oldRCWPassword.len)), []byte(C.GoStringN(newRCWPassword.data, newRCWPassword.len)), removeOldDir)
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

// GenTOTP returns:
// r0: err
//
// r1: token
//
//export GenTOTP
func GenTOTP(secret C.PascalString, unixTimestamp int64) (*C.char, C.PascalString) {
	token, err := core.GenTOTP(C.GoStringN(secret.data, secret.len), time.Unix(unixTimestamp, 0))
	if err != nil {
		return C.CString(err.Error()), C.PascalString{data: nil, len: -1}
	}
	return nil, getPascalString(token)
}

// GetOldEntryData returns:
// r0: err
//
// r1: lines
//
//export GetOldEntryData
func GetOldEntryData(realPath C.PascalString, field int, rcwPassword C.PascalString) (*C.char, C.PascalStringArray) {
	lines, err := core.GetOldEntryData(C.GoStringN(realPath.data, realPath.len), field, []byte(C.GoStringN(rcwPassword.data, rcwPassword.len)))
	if err != nil {
		return C.CString(err.Error()), C.PascalStringArray{}
	}
	return nil, getCPascalStringArrayFromStringSlice(lines)
}

// LibmuttonInit is a partial implementation that does not allow for client-specific config.
// Returns:
// r0: err
//
// Special requirement:
// An array of ordered responses to inputCB questions.
//
//export LibmuttonInit
func LibmuttonInit(inputCBResponses C.PascalStringArray, rcwPassword C.PascalString, appendMode bool, forceOfflineMode bool, deviceIDPrefix C.PascalString) *C.char {
	inputCBRespSlice := getStringSliceFromCPascalStringArray(inputCBResponses)
	supplyInputForInit := func(prompt string) string {
		if strings.HasPrefix(prompt, "Configure SSH settings (for synchronization)? (Y/n)") {
			return inputCBRespSlice[0]
		} else if strings.HasPrefix(prompt, "SSH private identity file path (falls back to") {
			return inputCBRespSlice[1]
		} else if strings.HasPrefix(prompt, "Is the identity file password-protected? (y/N)") {
			return inputCBRespSlice[2]
		} else if strings.HasPrefix(prompt, "Remote SSH username:") {
			return inputCBRespSlice[3]
		} else if strings.HasPrefix(prompt, "Remote SSH IP/domain:") {
			return inputCBRespSlice[4]
		} else if strings.HasPrefix(prompt, "Remote SSH port:") {
			return inputCBRespSlice[5]
		}
		return ""
	}

	if err := core.LibmuttonInit(supplyInputForInit, []byte(C.GoStringN(rcwPassword.data, rcwPassword.len)), appendMode, forceOfflineMode, C.GoStringN(deviceIDPrefix.data, deviceIDPrefix.len), nil); err != nil {
		return C.CString(err.Error())
	}
	return nil
}

// RCWSanityCheckGen returns:
// r0: err
//
//export RCWSanityCheckGen
func RCWSanityCheckGen(password C.PascalString) *C.char {
	err := core.RCWSanityCheckGen([]byte(C.GoStringN(password.data, password.len)))
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

// VerifyEntries returns:
// r0: err
//
//export VerifyEntries
func VerifyEntries(rcwPassword C.PascalString) *C.char {
	err := core.VerifyEntries([]byte(C.GoStringN(rcwPassword.data, rcwPassword.len)))
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

// WriteEntry returns:
// r0: err
//
//export WriteEntry
func WriteEntry(realPath C.PascalString, decSlice C.PascalStringArray, passwordIsNew bool, rcwPassword C.PascalString) *C.char {
	if err := core.WriteEntry(C.GoStringN(realPath.data, realPath.len), getStringSliceFromCPascalStringArray(decSlice), passwordIsNew, []byte(C.GoStringN(rcwPassword.data, rcwPassword.len))); err != nil {
		return C.CString(err.Error())
	}
	return nil
}
