# cmutton
Official C bindings for [libmutton](https://github.com/rwinkhart/libmutton).

This repository contains what is needed to generate a C library+headers for interfacing with libmutton
in native code. The bindings were written by hand and cover all libmutton functionality that is possible
to export to C.

# Usage
### Building
`go build -buildmode=c-archive`.
### Functions
All relevant exported libmutton functions have C functions of the same name.
Functions with multiple return values all have CGO-generated structs to store the return values.
These structs are named `<FunctionName>_return`, as per CGO.
All functions continue to perform the same basic operations, with a couple caveats:
- Anything that would normally return a Go error now returns a null-terminated *C.char containing the error string.
These are safe to null-terminate because their values are much more predictable than other strings.
Errors are always the *first* return value, so they can always be referenced with `<result>.r0`.
- Anything that would normally return a Go string or byte slice now returns a C.PascalString struct.
This is to avoid bugs with null-terminated strings.

Additionally, since comments tend to fall out of date, please rely on the [Go documentation for libmutton](https://pkg.go.dev/github.com/rwinkhart/libmutton). Documentation for the CGO bindings present in this repo only specify return values.
### Example (libmutton's global.DirInit)
Build the following example with `gcc <filename> ./cmutton.a`.
```c
#include <stdio.h>
#include "cmutton.h" // import cmutton

int main() {
    // use CGO-generated struct to get multiple return values
    struct DirInit_return result = DirInit(1);

    // familiar error handling pattern
    if (result.r0 != NULL) {
        printf("Error: %s\n", result.r0);
        exit(1); // result.r0 not freed since program exits
    }

    // print Pascal string using "%.*s" with printf and supplying both length and data
    printf("Old device ID: %.*s\n", result.r1.len, result.r1.data);

    // be sure to free the data!
    free(result.r1.data);
}
```

# Progress
- [ ] age
    - [ ] AllPasswordEntries(forceReage bool) error
    - [ ] Entry(vanityPath string, timestamp int64) error
    - [ ] TranslateAgeTimestamp(timestamp *int64) uint8
- [ ] clip
    - [ ] ClearArgument() error
    - [ ] ClearProcess(assignedContents string) error
    - [ ] CopyShortcut(realPath string, field int) error
    - [ ] CopyString(clearClipboardAutomatically bool, copySubject string) error
    - [ ] LaunchClearProcess(copySubject string)
    - [ ] TOTPCopier(secret string, errorChan chan<- error, done <-chan bool)
- [ ] config
    - [ ] Write(cfg *CfgT, appendMode bool) error
    - [ ] Load() (*CfgT, error)
- [ ] core
    - [ ] ClampTrailingWhitespace(note []string)
    - [ ] EntryAddPrecheck(realPath string) (uint8, error)
    - [ ] EntryIsNotEmpty(entryData []string) bool
    - [ ] EntryRefresh(oldRCWPassword, newRCWPassword []byte, removeOldDir bool) error
    - [ ] GenTOTP(secret string, time time.Time) (string, error)
    - [ ] GetOldEntryData(realPath string, field int) ([]string, error)
    - [ ] LibmuttonInit(inputCB func(prompt string) string, rcwPassword []byte, ...) error
    - [ ] RCWSanityCheckGen(password []byte) error
    - [ ] WriteEntry(realPath string, decSlice []string, passwordIsNew bool) error
- [ ] crypt
    - [ ] DecryptFileToSlice(realPath string) ([]string, error)
    - [ ] EncryptBytes(decBytes []byte) []byte
    - [ ] RCWDArgument()
    - [ ] VAR: Daemonize bool
    - [ ] VAR: RetryPassword bool
- [ ] global
    - [X] DirInit(preserveOldCfgDir bool) (string, error)
    - [ ] ~~GenDeviceIDList() ([]fs.DirEntry, error)~~ (not for use outside of libmutton)
    - [X] GetCurrentDeviceID() (string, error)
    - [X] GetRealAgePath(vanityPath string) string
    - [X] GetRealPath(vanityPath string) string
    - [ ] ~~GetSysProcAttr() *syscall.SysProcAttr~~ (not for use outside of libmutton)
    - [X] GetVanityPath(realPath string) string
    - [ ] VAR (CB func): GetPassword
- [ ] syncclient
    - [ ] AddFolderRemote(vanityPath string) error
    - [ ] GenDeviceID(oldDeviceID, prefix string) (string, string, bool, error)
    - [ ] GetSSHClient() (*ssh.Client, bool, *bool, *string, *string, error)
    - [ ] GetSSHOutput(sshClient *ssh.Client, cmd, stdin string) ([]byte, error)
    - [ ] RenameRemote(oldVanityPath, newVanityPath string) error
    - [ ] ShearRemote(vanityPath string, onlyShearAgeFile bool) error
    - [ ] RunJob() (*syncListsT, error)
- [ ] synccommon
    - [ ] AddFolderLocal(vanityPath string) error
    - [ ] RenameLocal(oldVanityPath, newVanityPath string) error
    - [ ] ShearAgeFileLocal(vanityPath string) error
    - [ ] ShearLocal(vanityPath, clientDeviceID string, onlyShearAgeFile bool) (string, bool, error)
    - [ ] WalkEntryDir() ([]string, []string, error)
    - [ ] GetAllEntryData() (EntryMapT, error)
- [ ] syncserver
    - [ ] GetRemoteDataFromServer(clientDeviceID string)
