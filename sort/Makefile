test:
	@ls -d *sort | awk '{print "./"$$0}' | xargs go test

generate:
	@go generate ./...

bench/ints:
	@go test -bench=BenchmarkInt -benchmem

bench/strings:
	@go test -bench=BenchmarkString -benchmem