package logit

import (
	"bytes"
	"errors"
	"testing"

	"github.com/go-logit/logit/core/appender"
)

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
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))

	options := Options()
	logger := NewLogger(
		options.WithDebugLevel(),
		options.WithAppender(appender.Json()),
		options.WithWriter(buffer),
		//options.WithPID(),
		//options.WithCaller(),
		//options.WithMsgKey("message"),
		options.WithTimeKey(""),
		//options.WithLevelKey("level"),
		//options.WithPIDKey("pid"),
		//options.WithFileKey("file"),
		//options.WithLineKey("line"),
		//options.WithTimeFormat("060102"),
	)

	logger.SetToGlobal()
	defer logger.Close()

	Debug("debug...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Any("any", map[string]interface{}{"a": 1, "b": "bbb"}).Log()
	Error("error...").Byte("b", 'a').Byte("es", '\n').Runes("words", []rune("我是中国人")).Error("err", errors.New("我是错误")).Log()
	Error("error with %d...", 666).String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
	Warn("\"warn\"...\r\b\t\n").Strings("s\tb\nd\b", []string{"abc\r", "efg\n"}).Log()
	Info("info...").Bools("bools", []bool{true, false}).Bytes("bytes", []byte{'\b', '\t', 'a', 'b', 'c', '"', '\n'}).Int16s("int16s", []int16{123, 4567, 8901}).Float32s("float32s", []float32{3.14, 6.18}).Log()

	logs := `{"log.level":"debug","log.msg":"debug...","trace":"xxx","id":123,"pi":3.14,"any":{"a":1,"b":"bbb"}}
{"log.level":"error","log.msg":"error...","b":"a","es":"\n","words":["我","是","中","国","人"],"err":"我是错误"}
{"log.level":"error","log.msg":"error with 666...","trace":"xxx","id":123,"pi":3.14}
{"log.level":"warn","log.msg":"\"warn\"...\r\b\t\n","s\tb\nd\b":["abc\r","efg\n"]}
{"log.level":"info","log.msg":"info...","bools":[true,false],"bytes":["\b","\t","a","b","c","\"","\n"],"int16s":[123,4567,8901],"float32s":[3.140000104904175,6.179999828338623]}
`

	output := buffer.String()
	if output != logs {
		t.Errorf("logs %s is wrong with %s", output, logs)
	}
}
