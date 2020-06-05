.PHONY: clean test run-tests tests/* release version readme

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

ok-macos.zip: version clean
	GOOS=darwin GOARCH=amd64 go build -o bin/ok
	zip $@ -r bin

ok-linux.zip: version clean
	GOOS=linux GOARCH=amd64 go build -o bin/ok
	zip $@ -r bin

ok-windows.zip: version clean
	GOOS=windows GOARCH=amd64 go build -o bin/ok
	zip $@ -r bin

release: ok-macos.zip ok-linux.zip ok-windows.zip

version: TRAVIS_TAG ?= $(shell git describe --tags --abbrev=0)
version: DATE := $(shell date +'%F')
version:
	sed -i '' "s/ok version unknown/ok version $(TRAVIS_TAG) $(DATE)/" cmd/version/main.go

readme:
	./gh-md-toc --insert README.md
	rm -f README.md.orig.* README.md.toc.*
