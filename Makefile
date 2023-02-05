.PHONY: buildandrun
BIN_FILE=run.out

build:
		@go build -o "${BIN_FILE}" main.go
		