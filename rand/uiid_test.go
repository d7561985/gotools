//
package rand

import (
	"testing"
	"fmt"
)

func TestRandStringBytes(t *testing.T) {
	a := RandStringBytes(33)
	b := RandStringBytes(33)
	if string(a) == string(b) {
		t.Errorf("% X == % X\n", a, b)
	}else {
		fmt.Printf("% X != % X\n", a, b)
		fmt.Printf("%s != %s\n", a, b)
	}
}

func BenchmarkUiid2(b *testing.B) {
	for i:=0; i < b.N;i++{
		Uiid()
	}
}

func BenchmarkRandStringBytes(b *testing.B) {
	for i:=0; i < b.N;i++{
		RandStringBytes(33)
	}
}