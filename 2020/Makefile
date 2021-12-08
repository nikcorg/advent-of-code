BINARY:=solve
COVERAGE_TMP:=c.out

all: test build run clean

build:
	@go build -o ${BINARY} *.go

test:
	@go test -coverprofile=${COVERAGE_TMP} ./...
	#@go tool cover -func=${COVERAGE_TMP}
	@unlink ${COVERAGE_TMP}

run: build
	@./${BINARY}

clean:
	@rm ${BINARY}

