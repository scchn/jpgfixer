package jpgfixer

import (
	"errors"
)

const (
	nmk = 4
	soi = 0xd8
	dht = 0xc4
	sos = 0xda
)

// Fix fixes a mjpeg frame by inserting the huffman table
func Fix(src []byte) ([]byte, error) {
	var (
		dst  []byte
		ht   bool
		data = append(src[:0], src...)
		mk   = make([]byte, nmk)
	)

	memcpy(mk, &data, 2)
	if mk[0] != 0xff || mk[1] != soi {
		return nil, errors.New("missing SOI")
	}
	dst = append(dst, mk[:2]...)

	for !ht {
		memcpy(mk, &data, nmk)
		if mk[0] != 0xff {
			return nil, errors.New("missing marker")
		}
		if mk[1] == dht {
			ht = true
		} else if mk[1] == sos {
			break
		}
		size := (int(mk[2]) << 8) | int(mk[3])
		dst = append(dst, mk...)
		l := size - 2
		tmp := make([]byte, l)
		memcpy(tmp, &data, l)
		dst = append(dst, tmp...)
	}

	if !ht {
		dst = append(dst, []byte(table)...)
		dst = append(dst, mk...)
	}

	dst = append(dst, data...)
	return dst, nil
}

func memcpy(dst []byte, src *[]byte, n int) int {
	if n > len(*src) {
		return 0
	}
	r := copy(dst, (*src)[:n])
	*src = (*src)[n:]
	return r
}
