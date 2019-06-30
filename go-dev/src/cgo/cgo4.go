package main

import "unsafe"

// #cgo CFLAGS: -I .
// #include "myprint.h"
import "C"

func main() {
	cs := C.CString("Hello from stdio\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}
