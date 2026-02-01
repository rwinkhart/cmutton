package main

// #include <stdlib.h>
// #include "types.h"
import "C"
import (
	"github.com/rwinkhart/libmutton/config"
)

// LoadOfficialCfg returns the currently set (libmutton official) values in the libmutton config file.
// Returns:
// r0: err
//
// r1: OfflineMode (bool)
//
// r2: SSHUser (string)
//
// r3: SSHIP (string)
//
// r4: SSHPort (string)
//
// r5: SSHEntryRootPath (string)
//
// r6: SSHAgeDirPath (string)
//
// r7: SSHKeyPath (string)
//
// r8: SSHKeyProtected (bool)
//
// r9: SSHIsWindows (bool)
//
//export LoadOfficialCfg
func LoadOfficialCfg() (*C.char, bool, C.PascalString, C.PascalString, C.PascalString, C.PascalString, C.PascalString, C.PascalString, bool, bool) {
	cfg, err := config.Load()
	if err != nil {
		return C.CString(err.Error()), false, getPascalString(""), getPascalString(""), getPascalString(""), getPascalString(""), getPascalString(""), getPascalString(""), false, false
	}
	return nil, safeBoolDeref(cfg.Libmutton.OfflineMode), getPascalString(safeStringDeref(cfg.Libmutton.SSHUser)), getPascalString(safeStringDeref(cfg.Libmutton.SSHIP)), getPascalString(safeStringDeref(cfg.Libmutton.SSHPort)), getPascalString(safeStringDeref(cfg.Libmutton.SSHEntryRootPath)), getPascalString(safeStringDeref(cfg.Libmutton.SSHAgeDirPath)), getPascalString(safeStringDeref(cfg.Libmutton.SSHKeyPath)), safeBoolDeref(cfg.Libmutton.SSHKeyProtected), safeBoolDeref(cfg.Libmutton.SSHIsWindows)
}

// WriteOfficialCfg writes the provided (libmutton official) values to the libmutton config file.
// Returns:
// r0: err
//
//export WriteOfficialCfg
func WriteOfficialCfg(appendMode bool, OfflineMode bool, SSHUser, SSHIP, SSHPort, SSHEntryRootPath, SSHAgeDirPath, SSHKeyPath C.PascalString, SSHKeyProtected, SSHIsWindows bool) *C.char {
	var cfg config.CfgT
	cfg.Libmutton.OfflineMode = &OfflineMode
	cfg.Libmutton.SSHUser = new(C.GoStringN(SSHUser.data, SSHUser.len))
	cfg.Libmutton.SSHIP = new(C.GoStringN(SSHIP.data, SSHIP.len))
	cfg.Libmutton.SSHPort = new(C.GoStringN(SSHPort.data, SSHPort.len))
	cfg.Libmutton.SSHEntryRootPath = new(C.GoStringN(SSHEntryRootPath.data, SSHEntryRootPath.len))
	cfg.Libmutton.SSHAgeDirPath = new(C.GoStringN(SSHAgeDirPath.data, SSHAgeDirPath.len))
	cfg.Libmutton.SSHKeyPath = new(C.GoStringN(SSHKeyPath.data, SSHKeyPath.len))
	cfg.Libmutton.SSHKeyProtected = &SSHKeyProtected
	cfg.Libmutton.SSHIsWindows = &SSHIsWindows
	err := config.Write(&cfg, appendMode)
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}
