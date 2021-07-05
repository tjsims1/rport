package enc

import (
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

const (
	bom0 = 0xef
	bom1 = 0xbb
	bom2 = 0xbf
)

func RemoveBomBytes(in string) string {
	bs_UTF8LE, _, _ := transform.Bytes(unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder(), []byte(in))
	bs_UTF8BE, _, _ := transform.Bytes(unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM).NewDecoder(), []byte(in))
	fmt.Println(string(bs_UTF8LE))
	fmt.Println(string(bs_UTF8BE))
	if len(in) >= 3 &&
		in[0] == bom0 &&
		in[1] == bom1 &&
		in[2] == bom2 {
		return in[3:]
	}
	return in
}
