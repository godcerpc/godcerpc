package ndr64

import (
	"testing"
)

func TestEncodeUint(t *testing.T) {
	testCases := []struct {
		in   interface{}
		want int
	}{
		{uint8(0x01), 1},
		{uint16(0x0123), 2},
		{uint32(0x01234567), 4},
		{uint64(0x0123456789abcdef), 8},
	}
	for _, tc := range testCases {
		gotBytes, gotNum := EncodeUint(tc.in)
		if len(gotBytes) != tc.want || gotNum != tc.want {
			t.Errorf("Wanted len %v, got: %v and %v", tc.want, len(gotBytes), gotNum)
		}
	}
}
