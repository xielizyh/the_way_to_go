package key

import (
	"crypto/md5"
	"strconv"
)

// GenKey 产生键值
func GenKey(n int, str string) string {
	hasher := md5.New()
	hasher.Write([]byte(strconv.Itoa(n)))
	hasher.Write([]byte(str))
	return string(hasher.Sum(nil))
}
