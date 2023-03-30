package logit

import (
	"bytes"
	"errors"
	"testing"

	"github.com/FishGoddess/logit/core/appender"
	"github.com/FishGoddess/logit/support/runtime"
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

	options := []Option{
		Options().WithDebugLevel(),
		Options().WithAppender(appender.Json()),
		Options().WithWriter(buffer),
		Options().WithCaller(),
		Options().WithMsgKey("msg"),
		Options().WithTimeKey(""),
		Options().WithLevelKey("level"),
		Options().WithPIDKey("pid"),
		Options().WithFileKey("file"),
		Options().WithLineKey("line"),
		Options().WithFuncKey("func"),
		Options().WithErrorKey("err"),
		Options().WithTimeFormat("060102"),
	}

	logger := NewLogger(options...).SetToGlobal()

	Debug("debug...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Any("any", map[string]interface{}{"a": 1, "b": "bbb"}).Log()
	Error(errors.New("我是错误"), "error...").Byte("b", 'a').Byte("es", '\n').Runes("words", []rune("我是中国人")).Log()
	Error(nil, "error with %d...", 666).String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
	Warn("\"warn\"...\r\b\t\n").Strings("s\tb\nd\b", []string{"abc\r", "efg\n"}).Log()
	Info("info...").Bools("bools", []bool{true, false}).Bytes("bytes", []byte{'\b', '\t', 'a', 'b', 'c', '"', '\n'}).Int16s("int16s", []int16{123, 4567, 8901}).Float32s("float32s", []float32{3.14, 6.18}).Log()
	logger.Close()

	file, _, _ := runtime.Caller(1)

	logs := `{"level":"debug","file":"` + file + `","line":44,"func":"github.com/FishGoddess/logit.TestGlobalLogger","msg":"debug...","trace":"xxx","id":123,"pi":3.14,"any":{"a":1,"b":"bbb"}}
{"level":"error","file":"` + file + `","line":45,"func":"github.com/FishGoddess/logit.TestGlobalLogger","msg":"error...","err":"我是错误","b":"a","es":"\n","words":["我","是","中","国","人"]}
{"level":"error","file":"` + file + `","line":46,"func":"github.com/FishGoddess/logit.TestGlobalLogger","msg":"error with 666...","err":null,"trace":"xxx","id":123,"pi":3.14}
{"level":"warn","file":"` + file + `","line":47,"func":"github.com/FishGoddess/logit.TestGlobalLogger","msg":"\"warn\"...\r\b\t\n","s\tb\nd\b":["abc\r","efg\n"]}
{"level":"info","file":"` + file + `","line":48,"func":"github.com/FishGoddess/logit.TestGlobalLogger","msg":"info...","bools":[true,false],"bytes":["\b","\t","a","b","c","\"","\n"],"int16s":[123,4567,8901],"float32s":[3.140000104904175,6.179999828338623]}
`

	output := buffer.String()
	if output != logs {
		t.Errorf("logs %s is wrong with %s", output, logs)
	}

	buffer = bytes.NewBuffer(make([]byte, 0, 1024))
	options = []Option{
		Options().WithDebugLevel(),
		Options().WithAppender(appender.Json()),
		Options().WithWriter(buffer),
		Options().WithCaller(),
		Options().WithMsgKey("msg"),
		Options().WithTimeKey(""),
		Options().WithLevelKey("level"),
		Options().WithPIDKey("pid"),
		Options().WithFileKey("file"),
		Options().WithLineKey(""),
		Options().WithFuncKey("func"),
		Options().WithErrorKey("err"),
		Options().WithTimeFormat("060102"),
	}

	logger = NewLogger(options...).SetToGlobal()
	defer logger.Close()

	Debug("debug...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Any("any", map[string]interface{}{"a": 1, "b": "bbb"}).Log()
	Error(errors.New("我是错误"), "error...").Byte("b", 'a').Byte("es", '\n').Runes("words", []rune("我是中国人")).Log()
	Error(nil, "error with %d...", 666).String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
	Warn("\"warn\"...\r\b\t\n").Strings("s\tb\nd\b", []string{"abc\r", "efg\n"}).Log()
	Info("info...").Bools("bools", []bool{true, false}).Bytes("bytes", []byte{'\b', '\t', 'a', 'b', 'c', '"', '\n'}).Int16s("int16s", []int16{123, 4567, 8901}).Float32s("float32s", []float32{3.14, 6.18}).Log()

	testBuffer := bytes.NewBuffer(make([]byte, 0, 1024))
	testLogger := NewLogger(append(options, Options().WithWriter(testBuffer))...)

	testLogger.Debug("debug...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Any("any", map[string]interface{}{"a": 1, "b": "bbb"}).Log()
	testLogger.Error(errors.New("我是错误"), "error...").Byte("b", 'a').Byte("es", '\n').Runes("words", []rune("我是中国人")).Log()
	testLogger.Error(nil, "error with %d...", 666).String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
	testLogger.Warn("\"warn\"...\r\b\t\n").Strings("s\tb\nd\b", []string{"abc\r", "efg\n"}).Log()
	testLogger.Info("info...").Bools("bools", []bool{true, false}).Bytes("bytes", []byte{'\b', '\t', 'a', 'b', 'c', '"', '\n'}).Int16s("int16s", []int16{123, 4567, 8901}).Float32s("float32s", []float32{3.14, 6.18}).Log()

	output = buffer.String()
	testOutput := testBuffer.String()

	if output != testOutput {
		t.Errorf("logs %s is wrong with %s", output, testOutput)
	}
}
