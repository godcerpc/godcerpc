package ndr64

import (
	"fmt"
)

// Unique pointers are transmitted as full pointers.
func EncodeUniqPtr(refId uint64, referent interface{}) (out []byte) {
	// TODO: Handle subsequent instances of the same pointer (i.e. aliasing)
	// TODO: Handle embedded pointers
	// FIXME: Ok this is a bad way of doing it. Yolo though for now

	switch referent.(type) {
	case uint8, uint16, uint32, uint64:
		out = append(out, EncodeUint(refId)...)
		out = append(out, EncodeUint(referent)...)
	// TODO: Probably refactor this so strings don't have to be wchar_t
	case string:
		out = append(out, EncodeUint(refId)...)
		out = append(out, EncodeWcharString(referent.(string))...)
	default:
		panic(fmt.Errorf("Type not supported for encoding: %T", referent))
	}
	return
}
