package main

// #include <stdlib.h>
// #include "types.h"
import "C"
import (
	"github.com/rwinkhart/libmutton/age"
)

// AgeEntry returns:
// err
//
//export AgeEntry
func AgeEntry(vanityPath C.PascalString, timestamp int64) *C.char {
	if err := age.Entry(C.GoStringN(vanityPath.data, vanityPath.len), timestamp); err != nil {
		return C.CString(err.Error())
	}
	return nil
}

// AgeAllPasswordEntries returns:
// err
//
//export AgeAllPasswordEntries
func AgeAllPasswordEntries(forceReage bool, rcwPassword C.PascalString) *C.char {
	if err := age.AllPasswordEntries(forceReage, []byte(C.GoStringN(rcwPassword.data, rcwPassword.len))); err != nil {
		return C.CString(err.Error())
	}
	return nil
}

// TranslateAgeTimestamp returns:
// ageStatusMagicNum (0 -> no age, 1 -> fresh, 2 -> expiring soon (within a month), 3 -> expired)
//
//export TranslateAgeTimestamp
func TranslateAgeTimestamp(timestamp *int64) uint8 {
	return age.TranslateAgeTimestamp(timestamp)
}
