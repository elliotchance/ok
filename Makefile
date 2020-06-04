.PHONY: clean test run-tests tests/*

ok:
	go build

clean:
	rm -f ok

test:
	go fmt ./...
	go vet ./...
	go test ./...
	make tests/*

run-tests: tests/*

tests/*: ok
	./ok run $@
