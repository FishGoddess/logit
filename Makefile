.PHONY: test bench benchfile fmt

all: test bench

test:
	go test -cover -count=1 -test.cpu=1 ./...

bench:
	go test -v ./_examples/performance_test.go -bench=^BenchmarkLogit -benchtime=1s

fmt:
	go fmt ./...