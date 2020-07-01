.PHONY: clean test run-tests tests/* release version readme ok lib-gen

ok: vm/lib.go
	go build

clean:
	rm -rf bin ok-macos.zip
	rm -f ok-linux.zip ok-windows.zip
	rm -f ok
	rm -f coverage.txt
	rm -f vm/lib.go

ci: clean test-fmt vet test-coverage run-tests check-readme check-doc

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

run-tests: tests/* run-lib-tests

tests/*: ok
	./ok run $@ > /tmp/stdout.txt
	diff $@/stdout.txt /tmp/stdout.txt

	./ok test $@

ok-macos.zip: version clean vm/lib.go
	GOOS=darwin GOARCH=amd64 go build -o bin/ok
	zip $@ -r bin

ok-linux.zip: version clean vm/lib.go
	GOOS=linux GOARCH=amd64 go build -o bin/ok
	zip $@ -r bin

ok-windows.zip: version clean vm/lib.go
	GOOS=windows GOARCH=amd64 go build -o bin/ok
	zip $@ -r bin

release: ok-macos.zip ok-linux.zip ok-windows.zip

version: TRAVIS_TAG ?= $(shell git describe --tags --abbrev=0)
version: DATE := $(shell date +'%F')
version:
	sed -i "s/ok version unknown/ok version $(TRAVIS_TAG) $(DATE)/" cmd/version/main.go

readme:
	./gh-md-toc --insert README.md
	sed -i.bak '/Added by:/d' README.md
	rm -f README.md.orig.* README.md.toc.*

check-readme:
	cp -f README.md README.md.bak2
	make readme
	diff README.md README.md.bak2
	rm -f README.md.bak README.md.bak2

lib-gen:
	go build ./cmd/lib-gen/

vm/lib.go: lib-gen
	./lib-gen

run-lib-tests:
	./ok test lib/math

check-doc:
	cp -f lib/math/README.md lib/math/README.md.bak2
	make doc
	diff lib/math/README.md lib/math/README.md.bak2
	rm -f lib/math/README.md.bak lib/math/README.md.bak2

doc: ok
	./ok doc lib/math > lib/math/README.md
