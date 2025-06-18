package kotlin

import (
	"fmt"
	"github.com/mtps/protoc-gen-cryptids/internal"
)

func AnyFromEAny(named string) string {
	v := fmt.Sprintf("%s.value.toByteArray()", named)
	decBs := DProvider(v)
	return fmt.Sprintf("com.google.protobuf.Any.parseFrom(%s)", decBs)
}

func EAnyFromAny(named string) string {
	v := fmt.Sprintf("%s.toByteArray()", named)
	encBs := EProvider(v)
	bstring := types.ByteStringCopyFrom(encBs)
	return fmt.Sprintf("%s.newBuilder().setValue(%s).build()", types.PackageFor(types.TypeEAny), bstring)
}
