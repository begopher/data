package complex

import(
	"hash/adler32"
)

type Strings []string

func (s Strings) Checksum() uint32 {
	hash := adler32.New()
	for _, line := range s {
		hash.Write([]byte(line))
	}
	return hash.Sum32()
}
