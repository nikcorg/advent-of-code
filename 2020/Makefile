BINARY:="solve"
TESTABLE_PACKAGES:=`find . -name "*_test.go" -exec dirname {} \; | sort | uniq`

all: test build run clean

build:
	go build -o ${BINARY} *.go

test:
	go test ${TESTABLE_PACKAGES}

run: build
	./${BINARY}

clean:
	rm ${BINARY}

