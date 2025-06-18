package kotlin

import (
	"fmt"
	"github.com/mtps/protoc-gen-cryptids/internal"
)

func StringFromEString(named string) string {
	rawBs := fmt.Sprintf("%s.value.toByteArray()", types.Uncapitalize(named))
	decBs := DProvider(rawBs)
	return fmt.Sprintf("String(%s)", decBs)
}

func EStringFromString(named string) string {
	rawEncBs := fmt.Sprintf("%s.toByteArray()", types.Uncapitalize(named))
	encBs := EProvider(rawEncBs)
	bs := types.ByteStringCopyFrom(encBs)
	return fmt.Sprintf("%s.newBuilder().setValue(%s).build()", types.PackageFor(types.TypeEString), bs)
}
