package ndr64

import (
	"fmt"
)

// Unique pointers are transmitted as full pointers
func EncodeUniqPtr(refId uint64, referent interface{}) []byte {
	switch referent.(type) {
	case uint32:
		// TODO: dostuff
		return []byte{}
	default:
		panic(fmt.Errorf("Type not supported for encoding: %T", referent))
	}
}
