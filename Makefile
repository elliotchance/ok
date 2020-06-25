.PHONY: clean test run-tests tests/* release version readme ok

ok:
	go build

clean:
	rm -rf bin ok-macos.zip
	rm -f ok-linux.zip ok-windows.zip
	rm -f ok
	rm -f coverage.txt

ci: clean test-fmt vet test-coverage run-tests

test:
	go test -race ./...

fmt:
	go fmt ./...

test-fmt:
	@echo "Checking go fmt..."
	exit $(shell go fmt ./... | wc -l)

vet:
	go vet ./...

test-coverage:
	echo "" > coverage.txt

	for d in $$(go list ./... | grep -v vendor); do \
		go test -race -coverprofile=profile.out -covermode=atomic $$d || exit 1; \
		if [ -f profile.out ]; then \
			cat profile.out >> coverage.txt; \
			rm profile.out; \
		fi \
	done

run-tests: tests/*

tests/*: ok
	./ok run $@ > /tmp/stdout.txt
	diff $@/stdout.txt /tmp/stdout.txt

	./ok test $@

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
	sed -i "s/ok version unknown/ok version $(TRAVIS_TAG) $(DATE)/" cmd/version/main.go

readme:
	./gh-md-toc --insert README.md
	rm -f README.md.orig.* README.md.toc.*
