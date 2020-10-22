package sodium

// #cgo pkg-config: libsodium
// #include <stdlib.h>
// #include <sodium.h>
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

func init() {
	if C.sodium_init() == -1 {
		fmt.Printf("libsodium init fail, terminate \n")
		os.Exit(1)
	}
	version := C.GoString(C.sodium_version_string())
	fmt.Printf("libsodium version is %s \n", version)
}

func MemZero(ptr unsafe.Pointer, size int) {
	defer fmt.Println("success do sodium_memzero")
	C.sodium_memzero(ptr, C.size_t(size))
}

func Malloc(size int) unsafe.Pointer {
	defer fmt.Println("success do sodium_malloc")
	ptr := C.sodium_malloc(C.size_t(size))
	if ptr == nil {
		panic("sodium_malloc failed")
	}
	return ptr
}

func Free(p unsafe.Pointer) {
	defer fmt.Println("success do sodium_free")
	C.sodium_free(p)
}

func MProtectNoAccess(p unsafe.Pointer) {
	if C.sodium_mprotect_noaccess(p) != 0 {
		fmt.Printf("libsodium mprotect no access fail, terminate \n")
		os.Exit(1)
	}
}

func MProtectReadOnly(p unsafe.Pointer) {
	if C.sodium_mprotect_readonly(p) != 0 {
		fmt.Printf("libsodium mprotect read only fail, terminate \n")
		os.Exit(1)
	}
}

func MProtectReadWrite(p unsafe.Pointer) {
	if C.sodium_mprotect_readwrite(p) != 0 {
		fmt.Printf("libsodium mprotect read write fail, terminate \n")
		os.Exit(1)
	}
}
