package ndr64

import (
	"bytes"
	"fmt"

	"testing"
)

func TestEncodeUniqPtr(t *testing.T) {
	testCases := []struct {
		refId    uint64
		referent interface{}
		want     []byte
	}{
		{1, uint32(0x01234567), []byte{1, 0, 0, 0, 0, 0, 0, 0, 103, 69, 35, 1}},
		{2, "localhost", []byte{2, 0, 0, 0, 0, 0, 0, 0, 0x0A, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0A, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6C, 0x00, 0x6F, 0x00, 0x63, 0x00, 0x61, 0x00, 0x6C, 0x00, 0x68, 0x00, 0x6F, 0x00, 0x73, 0x00, 0x74, 0x00, 0x00, 0x00}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%T:%v,refId:%v", tc.referent, tc.referent, tc.refId), func(t *testing.T) {
			gotBytes := EncodeUniqPtr(tc.refId, tc.referent)
			if bytes.Compare(tc.want, gotBytes) != 0 {
				t.Fail()
			}
		})
	}
}
