
build:
	go build

run: build
	go run main.go README.md
	cat README.md

all: build
