package jpgfixer

import (
	"errors"
)

const (
	nhdr = 4
	soi  = 0xd8
	dht  = 0xc4
	sos  = 0xda
)

// Fix fixes a mjpeg frame by inserting the huffman table
func Fix(imgData []byte) ([]byte, error) {
	var (
		dstData []byte
		data    = append(imgData[:0], imgData...)
		hasHdr  bool
		hdr     = make([]byte, nhdr)
	)

	memcpy(hdr, &data, 2)
	if hdr[0] != 0xff || hdr[1] != soi {
		return nil, errors.New("missing SOI")
	}
	dstData = append(dstData, hdr[:2]...)

	for !hasHdr {
		memcpy(hdr, &data, nhdr)
		if hdr[0] != 0xff {
			return nil, errors.New("missing marker")
		}
		if hdr[1] == dht {
			hasHdr = true
		} else if hdr[1] == sos {
			break
		}
		size := (int(hdr[2]) << 8) | int(hdr[3])
		dstData = append(dstData, hdr...)
		l := size - 2
		tmp := make([]byte, l)
		memcpy(tmp, &data, l)
		dstData = append(dstData, tmp...)
	}

	if !hasHdr {
		dstData = append(dstData, []byte(table)...)
		dstData = append(dstData, hdr...)
	}

	dstData = append(dstData, data...)
	return dstData, nil
}

func memcpy(dst []byte, src *[]byte, n int) int {
	if n > len(*src) {
		return 0
	}
	r := copy(dst, (*src)[:n])
	*src = (*src)[n:]
	return r
}
