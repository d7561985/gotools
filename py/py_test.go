//
package py

import "testing"

var env = []struct {
	t string
	res string

}{
	{string(`print(2**2)`), string("4")},
}

func Test_Run(t *testing.T) {

	var p Py
	for i := range env{
		res, err := p.Run(env[i].t)
		if err != nil{
			t.Errorf("%v\n", err)
		}
		if string(res[:len(res)-2]) != env[i].res{
			t.Errorf("res: %q should be %q\n", string(res[:len(res)-2]), env[i].res)
		}
	}
}

func BenchmarkPy_Run(b *testing.B) {
	for i:= 0; i < b.N; i++{
		var p Py
		for i := range env{
			_, err := p.Run(env[i].t)
			if err != nil{
				b.Errorf("%v\n", err)
			}
		}
	}
}