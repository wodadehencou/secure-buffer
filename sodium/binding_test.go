package sodium

import "testing"

func Test_Init(t *testing.T) {
}

func Test_Malloc(t *testing.T) {
	buf := Malloc(256)
	Free(buf)
}
