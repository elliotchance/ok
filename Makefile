.PHONY: clean test run-tests tests/* release

ok:
	go build

clean:
	rm -rf bin ok-macos.zip
	rm -f ok-linux.zip ok-windows.zip
	rm -f ok

test:
	go fmt ./...
	go vet ./...
	go test ./...
	make tests/*

run-tests: tests/*

tests/*: ok
	./ok run $@

ok-macos.zip: clean
	GOOS=darwin GOARCH=amd64 go build -o bin/ok
	zip $@ -r bin

ok-linux.zip: clean
	GOOS=linux GOARCH=amd64 go build -o bin/ok
	zip $@ -r bin

ok-windows.zip: clean
	GOOS=windows GOARCH=amd64 go build -o bin/ok
	zip $@ -r bin

release: ok-macos.zip ok-linux.zip ok-windows.zip
