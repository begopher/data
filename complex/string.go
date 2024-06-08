package complex

import(
	"hash/adler32"
)

type String string

func (s String) Checksum() uint32 {
	return adler32.Checksum([]byte(s))
}
