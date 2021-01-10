package key

/* // GenKey 产生键值
func GenKey(n int, str string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(strconv.Itoa(n)))
	hasher.Write([]byte(str))
	return hasher.Sum(nil)
}
*/

var keyChar = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// GenKey 产生键值
func GenKey(n int) string {
	if n == 0 {
		return string(keyChar[0])
	}
	l := len(keyChar)
	s := make([]byte, 20) // FIXME: will overflow. eventually.
	i := len(s)
	for n > 0 && i >= 0 {
		i--
		j := n % l
		n = (n - j) / l
		s[i] = keyChar[j]
	}
	return string(s[i:])
}
