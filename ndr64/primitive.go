package ndr64

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// https://pubs.opengroup.org/onlinepubs/9629399/chap14.htm#tagcjh_19_03_10
// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-dtyp/24637f2d-238b-4d22-b44d-fe54b024280c
// [MS-RPCE] 2.2.4 and 2.2.5
// https://docs.microsoft.com/en-us/windows/win32/midl/midl-predefined-and-base-types

// NDR64 always uses format label 0x10 (LE ASCII IEEE)

// https://pubs.opengroup.org/onlinepubs/9629399/chap4.htm#tagcjh_08_02_20
// my head hurts

// EncodeUint returns a byte slice containing the relevant NDR64 representation of any unsigned integer, and an int containing the number of bytes written
func EncodeUint(in interface{}) (out []byte, num int) {
	buf := new(bytes.Buffer)

	switch in.(type) {
	case uint8, uint16, uint32, uint64:
		binary.Write(buf, binary.LittleEndian, in)
		num = buf.Len()
		out = buf.Bytes()
	default:
		panic(fmt.Errorf("Type %T is not a uint", in))
	}
	return
}
