package main

/*
typedef struct {
    char* data;
    int len;
} PascalString;
*/
import "C"

// getPascalString returns a pascal
// string struct for the input Go string.
func getPascalString(goString string) C.PascalString {
	goStringBytes := []byte(goString)
	goStringPtr := C.CBytes(goStringBytes)
	return C.PascalString{
		data: (*C.char)(goStringPtr),
		len:  C.int(len(goStringBytes)),
	}
}

// getPascalStringFromBytes returns a pascal
// string struct for the input Go byte slice.
func getPascalStringFromBytes(goBytes []byte) C.PascalString {
	goBytesPtr := C.CBytes(goBytes)
	return C.PascalString{
		data: (*C.char)(goBytesPtr),
		len:  C.int(len(goBytes)),
	}
}
