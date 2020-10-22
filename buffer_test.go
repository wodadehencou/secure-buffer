package securebuffer

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"runtime"
	"testing"
)

func Test_NewBuffer(t *testing.T) {
	buf := New(256)
	for i, item := range buf.Bytes() {
		if item != 0xdb {
			t.Logf("item[%d] = %02x", i, item)
			t.Fail()
		}
	}

	buf = nil

	runtime.GC()

}

func Test_FromBytes(t *testing.T) {
	ori := make([]byte, 255)
	io.ReadFull(rand.Reader, ori)

	exp := make([]byte, len(ori))
	copy(exp, ori)

	buf := FromBytes(ori)

	// no open will panic
	buf.Open()
	defer buf.Close()

	act := buf.Bytes()
	if len(act) != len(exp) {
		t.Error("act length is not right")
	}
	for i := 0; i < len(exp); i++ {
		if act[i] != exp[i] {
			t.Error("data is not right")
		}
	}
	for i := 0; i < len(exp); i++ {
		if ori[i] != 0 {
			t.Error("ori buffer is not cleared")
		}
	}

	buf.Close()

	bs := buf.Bytes()
	t.Logf("length is %d", len(bs))
	// panic if read after Close
	// t.Logf("position 0 is %02x", bs[0])
}

func Test_Fault(t *testing.T) {
	t.Skip()
	ori := make([]byte, 10)
	buf := FromBytes(ori)
	buf.OpenRW()
	// buf.Open()
	bs := buf.Bytes()
	bs[0] = 0x55
	bs[9] = 0xff
	buf.Close()

	buf.Open()
	defer buf.Close()
	t.Log(hex.EncodeToString(buf.Bytes()))
}

// about 16k buffers can create
func Test_MostBuffer(t *testing.T) {
	t.Skip()
	ori := make([]byte, 16*1024)
	buffers := make([]*Buffer, 0)

	for i := 0; i < 1024*1024; i++ {
		t.Logf("No.%d buffer", i)
		io.ReadFull(rand.Reader, ori)
		buf := FromBytes(ori)
		buffers = append(buffers, buf)
	}

	t.Logf("total %d buffers", len(buffers))
}

func Test_MemoryLeak(t *testing.T) {
	// t.Skip()
	ori := make([]byte, 16*1024)

	for i := 0; i < 1024*1024; i++ {
		t.Logf("No.%d buffer", i)
		io.ReadFull(rand.Reader, ori)
		buf := FromBytes(ori)
		buf.Open()
		_ = buf.Bytes()
		buf.Close()
	}

}
