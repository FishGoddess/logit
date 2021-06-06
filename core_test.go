package logit

import (
	"os"
	"testing"
)

// go test -v -cover -run=^TestCore$
func TestCore(t *testing.T) {

	c := newCore(NewTextEncoder(TimeFormat), os.Stdout)
	c.SetLevel(WarnLevel)
	if c.Level() != WarnLevel {
		t.Fatalf("level %+v of core is wrong", c.Level())
	}

	c.SetNeedCaller(true)
	if c.NeedCaller() != true {
		t.Fatalf("needCaller %+v of core is wrong", c.NeedCaller())
	}
}
