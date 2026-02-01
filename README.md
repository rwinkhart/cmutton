# cmutton
Official C bindings for [libmutton](https://github.com/rwinkhart/libmutton).

This repository contains what is needed to generate a C library+headers for interfacing with libmutton
in native code. The bindings were written by hand and cover all libmutton functionality that is possible
to export to C.

# Usage
### Building
`go build -buildmode=c-archive`.
### Functions
All relevant exported libmutton functions have C counterparts with similar names.
Functions with multiple return values all have CGO-generated structs to store the return values.
These structs are named `<FunctionName>_return`, as per CGO.
All functions perform similar basic operations, with a couple caveats:
- Anything that would normally return a Go error now returns a null-terminated *C.char containing the error string.
These are safe to null-terminate because their values are much more predictable than other strings.
Errors are always the *first* return value, so they can always be referenced with `<result>.r0`.
- Many things that would normally require/return a Go string or byte slice now require/return a C.PascalString struct.
This is to avoid bugs with null-terminated strings.

Additionally, since comments tend to fall out of date, please rely on the [Go documentation for libmutton](https://pkg.go.dev/github.com/rwinkhart/libmutton). Documentation for the CGO bindings present in this repo only specify return values.
### Example (decrypting and printing lines of a libmutton entry)
Build the following example with `gcc <filename> ./cmutton.a`.
```c
#include <stdio.h>
#include <string.h>
#include "cmutton.h" // import cmutton

static void read_input(const char *prompt, char *buffer, size_t size) {
    printf("%s", prompt);
    fflush(stdout);
    if (fgets(buffer, size, stdin) == NULL) {
        fprintf(stderr, "Error reading user input\n");
        exit(1);
    }
    // remove trailing newline
    buffer[strlen(buffer)-1] = '\0';
}

int main() {
    char vanityPath[256];
    char password[256];
    read_input("Enter vanity path: ", vanityPath, sizeof(vanityPath));
    read_input("Enter password: ", password, sizeof(password));

    // use CGO-generated struct to get multiple return values
    struct DecryptFileToSlice_return result = DecryptFileToSlice(GetRealPath(vanityPath), GetPascalStringFromCString(password));

    // familiar error handling pattern
    if (result.r0 != NULL) {
        printf("Error: %s\n", result.r0);
        exit(1); // result.r0 not freed since program exits
    }

    for (int i = 0; i < result.r1.len; i++) {
        // print Pascal string using "%.*s" with printf and supplying both length and data
        printf("%.*s\n", result.r1.data[i].len, result.r1.data[i].data);
    }

    // be sure to free the data!
    FreeArray(result.r1);
}
```

# Progress
- [X] age
    - [X] AllPasswordEntries(forceReage bool, rcwPassword []byte) error
    - [X] Entry(vanityPath string, timestamp int64) error
    - [X] TranslateAgeTimestamp(timestamp *int64) uint8
- [X] clip
    - [X] ~~ClearArgument() error~~ (clipclear argument not supported)
    - [X] ClearProcess(assignedContents string) error 
    - [X] CopyShortcut(realPath string, field int) error
    - [X] CopyString(clearClipboardAutomatically bool, copySubject string) error
    - [X] ~~LaunchClearProcess(copySubject string)~~ (clipclear argument not supported)
    - [X] ~~TOTPCopier(secret string, errorChan chan<- error, done <-chan bool)~~ (not supported due to complexity)
- [X] config
    - [X] Write(cfg *CfgT, appendMode bool) error
    - [X] Load() (*CfgT, error)
- [ ] core
    - [ ] EntryAddPrecheck(realPath string) (uint8, error)
    - [ ] EntryIsNotEmpty(entryData []string) bool
    - [ ] EntryRefresh(oldRCWPassword, newRCWPassword []byte, removeOldDir bool) error
    - [ ] GenTOTP(secret string, time time.Time) (string, error)
    - [ ] GetOldEntryData(realPath string, field int, rcwPassword []byte) ([]string, error)
    - [ ] LibmuttonInit(inputCB func(prompt string) string, rcwPassword []byte, ...) error
    - [ ] RCWSanityCheckGen(password []byte) error
    - [ ] VerifyEntries(rcwPassword []byte) error
    - [ ] WriteEntry(realPath string, decSlice []string, passwordIsNew bool, rcwPassword []byte) error
- [X] crypt
    - [X] ~~VAR: RetryPassword bool~~ (RCWD not supported)
    - [X] DecryptFileToSlice(realPath string) ([]string, error)
    - [X] EncryptBytes(decBytes []byte) []byte
    - [X] ~~RCWDArgument()~~ (RCWD not supported)
- [X] global
    - [X] ~~VAR (CB func): GetPassword~~ (RCWD not supported)
    - [X] DirInit(preserveOldCfgDir bool) (string, error)
    - [X] ~~GenDeviceIDList() ([]fs.DirEntry, error)~~ (not for use outside of libmutton)
    - [X] GetCurrentDeviceID() (string, error)
    - [X] GetRealAgePath(vanityPath string) string
    - [X] GetRealPath(vanityPath string) string
    - [X] ~~GetSysProcAttr() *syscall.SysProcAttr~~ (not for use outside of libmutton)
    - [X] GetVanityPath(realPath string) string
- [ ] syncclient
    - [X] AddFolderRemote(vanityPath string) error
    - [X] GenDeviceID(oldDeviceID, prefix string) (string, string, bool, error)
    - [X] ~~GetSSHClient() (*ssh.Client, bool, *bool, *string, *string, error)~~ (not for use outside of libmutton)
    - [X] ~~GetSSHOutput(sshClient *ssh.Client, cmd, stdin string) ([]byte, error)~~ (not for use outside of libmutton)
    - [X] RenameRemote(oldVanityPath, newVanityPath string) error
    - [X] ShearRemote(vanityPath string, onlyShearAgeFile bool) error
    - [X] RunJob() (*syncListsT, error)
- [X] synccommon
    - [X] ~~AddFolderLocal(vanityPath string) error~~ (not for use outside of libmutton)
    - [X] ~~RenameLocal(oldVanityPath, newVanityPath string) error~~ (not for use outside of libmutton)
    - [X] ~~ShearAgeFileLocal(vanityPath string) error~~ (not for use outside of libmutton)
    - [X] ~~ShearLocal(vanityPath, clientDeviceID string, onlyShearAgeFile bool) (string, bool, error)~~ (not for use outside of libmutton)
    - [X] WalkEntryDir() ([]string, []string, error)
    - [X] GetAllEntryData() (EntryMapT, error)
- [X] syncserver
    - [X] ~~GetRemoteDataFromServer(clientDeviceID string)~~ (not for use outside of libmutton)
