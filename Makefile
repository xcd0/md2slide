
BIN=md2html_bin

build:
	go build -o $(BIN)

run: build
	./$(BIN) README.md

test: build
	go run test/test.go README.md

all: build


clean:
	rm *.html

