
build:
	go build

run: build
	go run main.go README.md

all: build
