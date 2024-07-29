test:
	go test ./...

run:
	go run examples/main.go

bench:
	go test -bench .
