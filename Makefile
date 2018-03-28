.PHONY: default install build test quicktest fmt vet lint

default: fmt vet lint build quicktest

install:
	go get -t -v ./...

build:
	go build -v ./...

test:
	go test -v -cover ./...

quicktest:
	go test ./...

fmt:
	@echo gofmt -l .
	@OUTPUT=$(gofmt -l . 2>&1); \
	    if [ "$$OUTPUT" ]; then \
		echo "gofmt must be run on the following files:"; \
		echo "$$OUTPUT"; \
		exit 1; \
	    fi

vet:
	go tool vet -atomic -bool -copylocks -nilfunc -printf -shadow -rangeloops -unreachable -unsafeptr -unusedresult .

lint:
	@echo golint ./...
	@OUTPUT=`golint ./... 2>&1`; \
	    if [ "$$OUTPUT" ]; then \
		echo "golint errors:"; \
		echo "$$OUTPUT"; \
		exit 1; \
	    fi

