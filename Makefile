all:
	go build ./cmd/...

gen:
	protoc --go_out=. --go_opt=paths=source_relative crypt/ext.proto crypt/types.proto
