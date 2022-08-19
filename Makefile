test:
	go test -v -cover ./...
bench:
	go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitLogger -benchtime=1s
benchfile:
	go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitFile -benchtime=1s