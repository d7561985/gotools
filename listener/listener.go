//
package listener

type Target func(interface{}) error

var L = Listener{make(map[string]Target)}

type Listenable interface {
	on(string string, f Target)
	off(string string)
	emit(string string)
}

type Listener struct {
	m map[string]Target
}

//! require to do Init when it is need.
func (l *Listener) Init() {
	l.m = make(map[string]Target)
}

func (l *Listener) On(s string, f Target) {
	_, ok := l.m[s]
	if ok {
		panic(s)
	}
	l.m[s] = f
}

func (l *Listener) Off(s string) {
	_, ok := l.m[s]
	if ok {
		delete(l.m, s)
	}
}

func (l *Listener) Emit(s string, i interface{}) error {
	c, ok := l.m[s]
	if !ok {
		return nil
	}
	return c(i)
}
