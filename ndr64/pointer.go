package ndr64

import (
	"fmt"
)

// Unique pointers are transmitted as full pointers.
func EncodeUniqPtr(refId uint64, referent interface{}) (out []byte) {
	// TODO: Handle subsequent instances of the same pointer (i.e. aliasing)
	// TODO: Handle embedded pointers
	switch referent.(type) {
	case uint8, uint16, uint32, uint64:
		out = append(out, EncodeUint(refId)...)
		out = append(out, EncodeUint(referent)...)
	default:
		panic(fmt.Errorf("Type not supported for encoding: %T", referent))
	}
	return
}
