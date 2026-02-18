# cmutton
Official C bindings for [libmutton](https://github.com/rwinkhart/libmutton).

This repository contains what is needed to generate a C library+headers for interfacing with libmutton
in native code. The bindings were written by hand and cover all libmutton functionality that is possible
to export to C.

# Usage
### Building
#### Generic
`CGO_ENABLED=1 go build -buildmode=c-archive`.
#### Apple Platforms (iOS/iOS Simulator/MacOS)
If building for Apple platforms, use the `Makefile` in the `xcode` directory.
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
