package ff

import (
	"hash/fnv"
)

func Hash9(s string) int64 {
	h := fnv.New32a()
	h.Write([]byte(s))
	hash := int64(h.Sum32())
	if hash < 100000000 {
		hash += 100000000
	}
	if hash > 999999999 {
		ex := hash - 999999999
		hash -= ex
	}
	return hash
}
