package kotlin

import (
	"fmt"
	"github.com/mtps/protoc-gen-cryptids/internal"
)

func ETimestampFromTimestamp(named string) string {
	tsBs := fmt.Sprintf(
		"java.nio.ByteBuffer.allocate(Int.SIZE_BYTES + Long.SIZE_BYTES).putLong(%s.seconds).putInt(%s.nanos).array()",
		named,
		named,
	)
	encBs := EProvider(tsBs)
	value := fmt.Sprintf("com.google.protobuf.ByteString.copyFrom(%s)", encBs)
	return fmt.Sprintf("%s.newBuilder().setValue(%s)", types.PackageFor(types.TypeETimestamp), value)
}

func TimestampFromETimestamp(named string) string {
	ts := DProvider(fmt.Sprintf("%s.value.toByteArray()", named))
	s := "com.google.protobuf.Timestamp.newBuilder().also { b -> val bb = java.nio.ByteBuffer.wrap(%s).order(java.nio.ByteOrder.BIG_ENDIAN); b.seconds = bb.long; b.nanos = bb.int; }.build()"
	return fmt.Sprintf(s, ts)
}
