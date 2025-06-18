package kotlin

import (
	"fmt"
	"github.com/mtps/protoc-gen-cryptids/internal"
)

func EBytesFromBytes(named string) string {
	encBs := types.ByteStringCopyFrom(EProvider(named))
	return fmt.Sprintf(
		"%s.newBuilder().setValue(%s).build()",
		types.PackageFor(types.TypeEBytes),
		encBs,
	)
}

func BytesFromEBytes(named string) string {
	bs := fmt.Sprintf("%s.value.toByteArray()", named)
	return DProvider(bs)
}
