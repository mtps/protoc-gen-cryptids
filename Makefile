all:
	go build ./cmd/...

gen:
	protoc --go_out=. --go_opt=paths=source_relative crypt/ext.proto crypt/types.proto

base: gen all
	mkdir -p ./test/src/generated/{java,kotlin}
	protoc --java_out=./test/src/generated/java \
               --kotlin_out=./test/src/generated/kotlin \
          test/test.proto crypt/types.proto

test: gen all
	mkdir -p ./test/src/generated/{java,kotlin}
	protoc --java_out=./test/src/generated/java \
               --kotlin_out=./test/src/generated/kotlin \
               --cryptids_out=paths=source_relative:./test/src/generated/java \
          test/test.proto crypt/types.proto
 
