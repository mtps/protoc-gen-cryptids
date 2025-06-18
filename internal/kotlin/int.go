package kotlin

import (
	"fmt"
	"github.com/mtps/protoc-gen-cryptids/internal"
)

func IntFromEInt(named string) string {
	v := fmt.Sprintf("%s.value.toByteArray()", named)
	decBs := DProvider(v)
	return fmt.Sprintf("java.nio.ByteBuffer.wrap(%s).order(java.nio.ByteOrder.BIG_ENDIAN).int", decBs)
}

func EIntFromInt(named string) string {
	bs := fmt.Sprintf("java.nio.ByteBuffer.allocate(Int.SIZE_BYTES).putInt(%s).order(java.nio.ByteOrder.BIG_ENDIAN).array()", named)
	encBs := EProvider(bs)
	bstring := types.ByteStringCopyFrom(encBs)
	return fmt.Sprintf("%s.newBuilder().setValue(%s).build()", types.PackageFor(types.TypeEInt), bstring)
}
