package main

import "C"
import (
	"path/filepath"

	"github.com/rwinkhart/go-boilerplate/back"
	"github.com/rwinkhart/libmutton/global"
)

func init() {
	// get the app's Library/Application Support directory (private, not backed up by iCloud)
	back.Home = filepath.Join(back.Home, "Library", "Application Support", "Passture")

	// change path variables to account for iOS app sandbox
	global.EntryRoot = back.Home + "/entries"             // Path to libmutton entry directory
	global.CfgDir = back.Home + "/config"                 // Path to libmutton configuration directory
	global.CfgPath = global.CfgDir + "/libmuttoncfg.json" // Path to libmutton configuration file
	global.AgeDir = global.CfgDir + "/age"                // Path to libmutton password age directory
	global.RootLength = len(global.EntryRoot)
}
