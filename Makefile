.PHONY: buildandrun
BIN_FILE=run.out

buildandrun:
		@go build -o "${BIN_FILE}" main.go
		./"${BIN_FILE}"