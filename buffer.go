package securebuffer

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/wodadehencou/securebuffer/sodium"
)

type slice struct {
	addr uintptr
	len  int
	cap  int
}

type Buffer struct {
	sl slice
}

func New(size int) *Buffer {
	buf := new(Buffer)
	buf.sl.addr = uintptr(sodium.Malloc(size))
	buf.sl.len = size
	buf.sl.cap = size
	runtime.SetFinalizer(buf, bufferFinalizer)
	return buf
}

func FromBytes(bs []byte) *Buffer {
	b := New(len(bs))
	copy(pointerToSlice(&b.sl), bs)
	defer clearBytes(bs)
	sodium.MProtectNoAccess(unsafe.Pointer(b.sl.addr))
	return b
}

func (b *Buffer) Open() {
	sodium.MProtectReadOnly(unsafe.Pointer(b.sl.addr))
}

func (b *Buffer) OpenRW() {
	sodium.MProtectReadWrite(unsafe.Pointer(b.sl.addr))
}

func (b *Buffer) Close() {
	sodium.MProtectNoAccess(unsafe.Pointer(b.sl.addr))
}

func (b *Buffer) Bytes() []byte {
	return pointerToSlice(&b.sl)
}

func pointerToSlice(sl *slice) []byte {
	return *(*[]byte)(unsafe.Pointer(sl))
}

func clearBytes(bs []byte) {
	sodium.MemZero(unsafe.Pointer(&bs[0]), len(bs))
}

func bufferFinalizer(buf *Buffer) {
	defer fmt.Println("do finalizer of buffer")
	sodium.Free(unsafe.Pointer(buf.sl.addr))
}
