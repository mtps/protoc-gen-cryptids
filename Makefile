all:
	go build ./cmd/...

gen:
	protoc --go_out=. --go_opt=paths=source_relative crypt/ext.proto crypt/types.proto

test: gen all
	mkdir -p ./src
	protoc --kotlin_out=./src --cryptids_out=paths=source_relative:./src test/test.proto
 
