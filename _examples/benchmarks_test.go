// Author: fish
// Email: fishinlove@163.com
// Created at 2020/03/02 21:07:01

package main

import (
    "log"
    "testing"

    "github.com/FishGoddess/logit"
    //"github.com/kataras/golog"
    //"github.com/sirupsen/logrus"
)

// 仅为测试用，不输出任何信息的写出器
type nopWriter struct{}

func (nw *nopWriter) Write(p []byte) (n int, err error) {
    return 0, nil
}

// 测试 logit 日志记录器的速度
func BenchmarkLogitLogger(b *testing.B) {

    // 测试用的日志记录器
    logger := logit.NewLogger(&nopWriter{}, logit.DebugLevel)

    // 测试用的日志任务
    logTask := func() {
        logger.Debug("debug...")
        logger.Info("info...")
        logger.Warn("warn...")
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
    logger := log.New(&nopWriter{}, "", log.LstdFlags)

    // 测试用的日志任务
    logTask := func() {
        logger.Println("debug...")
        logger.Println("info...")
        logger.Println("warn...")
        logger.Println("error...")
    }

    // 开始性能测试
    b.ReportAllocs()
    b.StartTimer()

    for i := 0; i < b.N; i++ {
        logTask()
    }
}

// 测试 golog 日志记录器的速度
//func BenchmarkGologLogger(b *testing.B) {
//
//    golog.SetOutput(&nopWriter{})
//    golog.SetLevel("debug")
//    golog.SetTimeFormat("")
//
//    // 测试用的日志任务
//    logTask := func() {
//        golog.Debug("debug...")
//        golog.Info("info...")
//        golog.Warn("warn...")
//        golog.Error("error...")
//    }
//
//    // 开始性能测试
//    b.ReportAllocs()
//    b.StartTimer()
//
//    for i := 0; i < b.N; i++ {
//        logTask()
//    }
//}

// 测试 logrus 日志记录器的速度
//func BenchmarkLogrusLogger(b *testing.B) {
//
//    logrus.SetOutput(&nopWriter{})
//    logrus.SetLevel(logrus.DebugLevel)
//
//    // 测试用的日志任务
//    logTask := func() {
//        logrus.Debug("debug...")
//        logrus.Info("info...")
//        logrus.Warn("warn...")
//        logrus.Error("error...")
//    }
//
//    b.ReportAllocs()
//    b.StartTimer()
//
//    for i := 0; i < b.N; i++ {
//        logTask()
//    }
//}