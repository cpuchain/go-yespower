package yespower

/*
#cgo CFLAGS: -std=gnu99 -march=native
#include <yespower/sha256.c>
#include <yespower/yespower-opt.c>
#include <yespower/yespower.c>
*/
import "C"

import (
	"unsafe"
)

func Hash(input []byte, per string) []byte {
	var in unsafe.Pointer = C.CBytes(input)
	var cPer unsafe.Pointer = unsafe.Pointer(C.CString(per))
	var out unsafe.Pointer = C.malloc(32)

	C.yespower_hash((*C.char)(in), C.uint(len(input)), (*C.char)(cPer), C.uint(len(per)), (*C.char)(out))

	hashed := C.GoBytes(out, 32)

	C.free(in)
	C.free(cPer)
	C.free(out)

	return hashed
}