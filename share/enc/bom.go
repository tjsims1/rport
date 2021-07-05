package enc

const (
	bom0 = 0xef
	bom1 = 0xbb
	bom2 = 0xbf
)

func RemoveBomBytes(in string) string {
	if len(in) >= 3 &&
		in[0] == bom0 &&
		in[1] == bom1 &&
		in[2] == bom2 {
		return in[3:]
	}
	return in
}
