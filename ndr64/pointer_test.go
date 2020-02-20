package ndr64

import (
	"bytes"

	"testing"
)

func TestEncodeUniqPtr(t *testing.T) {
	wantBytes := []byte{1, 0, 0, 0, 0, 0, 0, 0, 103, 69, 35, 1}
	gotBytes := EncodeUniqPtr(uint64(1), uint32(0x01234567))
	if bytes.Compare(wantBytes, gotBytes) != 0 {
		t.Fail()
	}
}
