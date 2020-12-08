BINARY:="solve"

all: test build run clean

build:
	go build -o ${BINARY} *.go

test:
	go test ./...

run: build
	./${BINARY}

clean:
	rm ${BINARY}

