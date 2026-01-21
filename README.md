# cmutton
Official C bindings for [libmutton](https://github.com/rwinkhart/libmutton).

# Usage
Build the C library + headers w/`go build -buildmode=c-archive`.

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
- [ ] global
    - [X] DirInit(preserveOldCfgDir bool) (string, error)
    - [ ] ~~GenDeviceIDList() ([]fs.DirEntry, error)~~
    - [ ] GetCurrentDeviceID() (string, error)
    - [ ] GetRealAgePath(vanityPath string) string
    - [ ] GetRealPath(vanityPath string) string
    - [ ] GetSysProcAttr() *syscall.SysProcAttr
    - [ ] GetVanityPath(realPath string) string
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
