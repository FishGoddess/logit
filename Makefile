.PHONY: test bench benchfile fmt

all: test bench

test:
	go test -cover ./...

bench:
	go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitLogger -benchtime=1s

benchfile:
	go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitFile -benchtime=1s

fmt:
	go fmt ./...