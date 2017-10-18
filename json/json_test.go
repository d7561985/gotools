// using reflection we can't do much more faster then original encode/json :(
package json

import "testing"

type Third struct {
	F string `json:"fa"`
}
type Tmp struct {
	Val  []string `json:"val"`
	//Val2 Third    `json:"vala"`
}

type CLogin struct {
	Cmd       string `json:"command"`
	Aaam      Tmp    `json:"tmp"`
	RequestId string `json:"request_id"`
	Token     string `json:"token"`
	Language  string `json:"language"`
}

var p = []struct {
	b []byte
}{
	{[]byte(`{"command":"login", "tmp": {"val":["111", "222"], "vala":{"fa":"aaa"}}, "request_id":"97144bf341554b05b12a8e11594a0b70","token":"USD:zodiac-d75619852@gmail.com","language":"en"}`)},
	{[]byte(`{"command":"login", "tmp": {"val":["111", "222"]}, "request_id":"97144bf341554b05b12a8e11594a0b70","token":"USD:zodiac-d75619852@gmail.com","language":"en"}`)},
}

func TestUnmarshal(t *testing.T) {
	var value CLogin
	err := Unmarshal(p[0].b, &value)
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkGetLogin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var value CLogin
		Unmarshal(p[0].b, &value)
	}
}
