//
package rand

import (
	r "math/rand"
	"time"
)

func init() {
	r.Seed(time.Now().UnixNano())
}

const tpl = `xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx`
//BenchmarkUiid2-8                 1000000              1035 ns/op              48 B/op          1 allocs/op
func Uiid() []byte {
	var res = make([]byte, 0, len(tpl))

	for i := 0; i < len(tpl); i++ {
		switch tpl[i] {
		case 'x':
			res = append(res, byte(48+r.Int63n(10))) //ancii digits
		case 'y':
			res = append(res, byte(97+r.Int63n(26))) //ancii words.
		default:
			res = append(res, tpl[i])
		}
	}
	return res
}

const letters = "abcdefghijklmnopqrstuvwxyz1234567890"
//BenchmarkRandStringBytes-8       2000000               859 ns/op              48 B/op          1 allocs/op
func RandStringBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return b
}
