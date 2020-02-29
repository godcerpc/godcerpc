package ndr64

import (
	// "fmt"
	"golang.org/x/text/encoding/unicode"
)

func EncodeWcharString(in string) (out []byte) {
	// wchar_t designates a wide character type. It is treated as an unsigned short by using the rules for an unsigned short
	// A string attribute can be applied to a pointer or array of type wchar_t. This indicates a string of wide characters, as specified in [C706] section 14.3.4.

	// If someone can tell me when a string should be conformant or conformant and variant that would pwn
	// Going with conformant and variant for now

	// The first integer gives the maximum number of elements in the string, including the terminator
	// The second integer gives the offset from the first index of the string to the first index of the actual subset being passed
	// The third integer gives the actual number of elements being passed, including the terminator
	// :shrug: just set them to string len and 0 hey?
	maxCount, actualCount := uint64(len(in)+1), uint64(len(in)+1)
	offset := uint64(0)

	out = append(out, EncodeUint(maxCount)...)
	out = append(out, EncodeUint(offset)...)
	out = append(out, EncodeUint(actualCount)...)

	encoding := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	encoder := encoding.NewEncoder()
	toAppend, err := encoder.Bytes([]byte(in))
	if err != nil {
		panic(err)
	}
	out = append(out, toAppend...)

	// The terminator for a wide character string is two octets of zero
	out = append(out, []byte{0x00, 0x00}...)
	return
}
