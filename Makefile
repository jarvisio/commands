test:
	go test

check:
	go vet *.go
	golint *.go

coverage:
	go test -coverprofile=coverage.out -covermode=count && go tool cover -html=coverage.out && unlink coverage.out

build:
	go build colorcat.go

install: build
	cp colorcat ~/bin/

install_devtools:
	go get code.google.com/p/go.tools/cmd/vet
	go get github.com/golang/lint/golint
	go get code.google.com/p/go.tools/cmd/cover
	go get code.google.com/p/go.tools/cmd/godoc

