.PHONY: clean test run-tests tests/* release version ok lib-gen

ok: version fs/lib.go
	go build

	# Restore the original main.go so that git does not track the changes.
	mv -f cmd/version/main.go.bak cmd/version/main.go

clean:
	rm -rf bin ok-macos.zip
	rm -f ok-linux.zip ok-windows.zip
	rm -f ok
	rm -f coverage.txt

ci: clean test-fmt vet test-coverage run-tests check-doc check-stdlib

test:
	go test -race ./...

fmt:
	go fmt ./...

test-fmt:
	@echo "Checking go fmt..."
	exit $(shell go fmt ./... | wc -l)

vet:
	# vet doesn't have a way to exclude files, so we have to do this hacky
	# workaround for now.
	mv -f fs/lib.go fs/lib.notgo
	go vet ./...
	mv -f fs/lib.notgo fs/lib.go

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

	@if [ -f "$@/stdout-ignored.txt" ]; then \
		echo "Stdout ignored."; \
	else \
		diff $@/stdout.txt /tmp/stdout.txt; \
	fi

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
	sed -i.bak "s/ok version unknown/ok version $(TRAVIS_TAG) $(DATE)/" cmd/version/main.go

lib-gen:
	go build ./cmd/lib-gen/

fs/lib.go: lib-gen
	./lib-gen
	go fmt fs/lib.go

run-lib-tests:
	for d in $(shell ls -d lib/*/); do \
        ./ok test $$d || exit 1 ; \
    done

check-doc:
	for d in $(shell ls -d lib/*/); do \
    	cp -f $$d/README.md $$d/README.md.bak2 ; \
    done

	make doc

	for d in $(shell ls -d lib/*/); do \
    	diff $$d/README.md $$d/README.md.bak2 || exit 1 ; \
    	rm -f $$d/README.md.bak $$d/README.md.bak2 ; \
    done

check-stdlib:
	mv -f fs/lib.go fs/lib.go.bak
	make fs/lib.go
	diff fs/lib.go fs/lib.go.bak

doc: ok
	for d in $(shell ls -d lib/*/); do \
		cd $$d && ../../ok doc > README.md && cd - ; \
    done

install: ok
	cp ok /usr/local/bin
