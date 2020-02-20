package ndr64

import (
	"bytes"

	"testing"
)

func TestEncodeUint(t *testing.T) {
	testCases := []struct {
		in   interface{}
		want []byte
	}{
		{uint8(0x01), []byte{0x01}},
		{uint16(0x0123), []byte{0x23, 0x01}},
		{uint32(0x01234567), []byte{0x67, 0x45, 0x23, 0x01}},
		{uint64(0x0123456789abcdef), []byte{0xef, 0xcd, 0xab, 0x89, 0x67, 0x45, 0x23, 0x01}},
	}
	for _, tc := range testCases {
		gotBytes := EncodeUint(tc.in)
		if bytes.Compare(tc.want, gotBytes) != 0 {
			t.Fail()
		}
	}
}
