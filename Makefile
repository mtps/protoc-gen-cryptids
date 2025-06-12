all:
	go build ./cmd/...

gen:
	protoc --go_out=. --go_opt=paths=source_relative crypt/ext.proto crypt/types.proto

test: gen all
	mkdir -p ./src/generated/kotlin
	protoc --kotlin_out=./src/generated/kotlin --cryptids_out=paths=source_relative:./src/generated/kotlin test/test.proto
 
