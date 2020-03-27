package jpgfixer

import (
	"bytes"
	"testing"
)

func Test_Memcpy(t *testing.T) {
	var (
		rl  = 4
		o   = []byte("12345678")
		dst = make([]byte, rl)
		src = []byte("12345678")
	)
	for len(src) != 0 {
		memcpy(dst, &src, rl)
		read := o[:rl]
		if bytes.Compare(dst, read) != 0 {
			t.Fatalf("Expect %s got %s", read, dst)
		}
		rest := o[rl:]
		if bytes.Compare(src, rest) != 0 {
			t.Fatalf("Expect %s got %s", rest, src)
		}
		o = rest
	}
}
