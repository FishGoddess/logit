package logit

import "testing"

// go test -v -cover -run=^TestSetGlobal$
func TestSetGlobal(t *testing.T) {
	logger := NewLogger()
	SetGlobal(logger)

	if globalLogger != logger {
		t.Errorf("globalLogger %p != logger %p", globalLogger, logger)
	}
}

// go test -v -cover -run=^TestGlobalLogger$
func TestGlobalLogger(t *testing.T) {

}
