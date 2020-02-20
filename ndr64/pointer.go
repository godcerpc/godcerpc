package ndr64

import (
	"fmt"
)

// Unique pointers are transmitted as full pointers.
func EncodeUniqPtr(refId uint64, referent interface{}) (out []byte, num int) {
	// TODO: Handle subsequent instances of the same pointer (i.e. aliasing)
	// TODO: Handle embedded pointers
	switch referent.(type) {
	case uint8, uint16, uint32, uint64:
		toAppend, toAdd := EncodeUint(refId)
		out = append(out, toAppend...)
		num += toAdd

		toAppend, toAdd = EncodeUint(referent)
		out = append(out, toAppend...)
		num += toAdd
	default:
		panic(fmt.Errorf("Type not supported for encoding: %T", referent))
	}
	return
}
