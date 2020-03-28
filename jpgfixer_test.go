package jpgfixer

import (
	"bytes"
	"testing"
)

func Test_Memcpy(t *testing.T) {
	var (
		offset = 0
		src    = []byte("12345678")
		o      = append([]byte{}, src...)
		dst    = make([]byte, 4)
	)
	for len(src) != 0 {
		read := memcpy(dst, &src, len(dst))
		if bytes.Compare(dst, o[offset:offset+read]) != 0 {
			t.Fatalf("dst expect %s, got %s", o[offset:offset+read], dst)
		}
		offset += read
		if bytes.Compare(src, o[offset:]) != 0 {
			t.Fatalf("src expect %s, got %s", o[offset:], src)
		}
	}
}
