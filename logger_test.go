// Author: fish
// Email: fishinlove@163.com
// Created at 2020/02/29 16:41:41

package logit

import (
	"log"
	"os"
	"testing"
)

// 测试日志记录器的 Debug 方法
func TestLoggerDebug(t *testing.T) {
	logger := NewLogger(os.Stdout, DebugLevel)
	logger.Debug("这是 debug 信息。。。")
}

// 测试日志记录器的 Debug 方法
func TestLoggerInfo(t *testing.T) {
	logger := NewLogger(os.Stdout, DebugLevel)
	logger.Info("这是 info 信息。。。")
}

// 测试日志记录器的 Debug 方法
func TestLoggerWarning(t *testing.T) {
	logger := NewLogger(os.Stdout, DebugLevel)
	logger.Warning("这是 warning 信息。。。")
}

// 测试日志记录器的 Debug 方法
func TestLoggerError(t *testing.T) {
	logger := NewLogger(os.Stdout, DebugLevel)
	logger.Error("这是 error 信息。。。")
}

// 测试调整日志记录器为运行状态的方法
func TestLoggerEnable(t *testing.T) {
	logger := NewLogger(os.Stdout, DebugLevel)
	logger.Debug("1. 这是 debug 信息。。。")
	logger.Disable()
	logger.Debug("2. 这是 debug 信息。。。")
	logger.Enable()
	logger.Debug("3. 这是 debug 信息。。。")
}

// 测试日志记录器的级别控制是否可用
func TestLoggerLevel(t *testing.T) {
	logger := NewLogger(os.Stdout, WarningLevel)
	logger.Info("这条 info 级别的内容可以显示吗？")
	logger.Error("这条 error 级别的内容可以显示吗？")
	logger.Warning("这条 warning 级别的内容可以显示吗？")
}

// ==============================================================
// Benchmarks
// ==============================================================

// 仅为测试用，不输出任何信息的写出器
type nopWriter struct{}

func (w *nopWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

// 测试 logit 日志记录器的速度
func BenchmarkLogitLogger(b *testing.B) {

	// 测试用的日志记录器
	logger := NewLogger(&nopWriter{}, DebugLevel)

	// 测试用的日志任务
	logTask := func() {
		logger.Debug("debug...")
		logger.Info("info...")
		logger.Warning("warning...")
		logger.Error("error...")
	}

	// 开始性能测试
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logTask()
	}
}

// 测试标准库 log.Logger 日志记录器的速度
func BenchmarkLogLogger(b *testing.B) {

	// 测试用的日志记录器
	logger := log.New(&nopWriter{}, "", log.LstdFlags|log.Lshortfile)

	// 测试用的日志任务
	logTask := func() {
		logger.Println("debug...")
		logger.Println("info...")
		logger.Println("warning...")
		logger.Println("error...")
	}

	// 开始性能测试
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logTask()
	}
}
