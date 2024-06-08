package complex

import(
	"hash/adler32"
)

type Bytes []byte

func (b Bytes) Checksum() uint32 {
	return adler32.Checksum(b)
}
