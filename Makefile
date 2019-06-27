ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
APP_NAME:=tl

run:
	go run $(ROOT_DIR)/main.go

build:
	@ go build -mod=vendor -o $(APP_NAME) main.go

test:
	go test $(ROOT_DIR)/translator

lint:
	golint $(ROOT_DIR)

linux-amd64:
	@ GOOS=linux GOARCH=amd64 go build -o $(APP_NAME)-linux-amd64 main.go

linux-armv8:
	@ GOOS=linux GOARCH=arm64 go build -o $(APP_NAME)-linux-armv8 main.go

linux-armv7:
	@ GOOS=linux GOARCH=arm GOARM=7 go build -o $(APP_NAME)-linux-armv7 main.go

linux-armv6:
	@ GOOS=linux GOARCH=arm GOARM=6 go build -o $(APP_NAME)-linux-armv6 main.go

linux-armv5:
	@ GOOS=linux GOARCH=arm GOARM=5 go build -o $(APP_NAME)-linux-armv5 main.go

darwin-amd64:
	@ GOOS=darwin GOARCH=amd64 go build -o $(APP_NAME)-darwin-amd64 main.go

freebsd-amd64:
	@ GOOS=freebsd GOARCH=amd64 go build -o $(APP_NAME)-freebsd-amd64 main.go

openbsd-amd64:
	@ GOOS=openbsd GOARCH=amd64 go build -o $(APP_NAME)-openbsd-amd64 main.go

windows-amd64:
	@ GOOS=windows GOARCH=amd64 go build -o $(APP_NAME)-windows-amd64 main.go

build-all: linux-amd64 linux-armv8 linux-armv7 linux-armv6 linux-armv5 darwin-amd64 freebsd-amd64 openbsd-amd64 windows-amd64

clean:
	@ rm -f $(ROOT_DIR)/$(APP_NAME)-* $(APP_NAME)
