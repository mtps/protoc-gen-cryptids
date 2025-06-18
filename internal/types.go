package types

import (
	"fmt"
	crypt "github.com/mtps/protoc-gen-cryptids/crypt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"regexp"
	"strings"
)

// TODO - Better package detection of internal types.
func PackageFor(descriptor protoreflect.Descriptor) string {
	return "com.github.mtps.protobuf.crypt.CryptProto." + string(descriptor.Name())
}

var TypeEString = (&crypt.EString{}).ProtoReflect().Descriptor()
var TypeEBytes = (&crypt.EBytes{}).ProtoReflect().Descriptor()
var TypeETimestamp = (&crypt.ETimestamp{}).ProtoReflect().Descriptor()
var TypeEInt = (&crypt.EInt{}).ProtoReflect().Descriptor()
var TypeEAny = (&crypt.EAny{}).ProtoReflect().Descriptor()

func Capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func Uncapitalize(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

func ByteStringCopyFrom(value string) string {
	return fmt.Sprintf("com.google.protobuf.ByteString.copyFrom(%s)", value)
}

var snakeCaseRx = regexp.MustCompile(`_([a-z])`)

func SnakeToCamel(s string) string {
	return snakeCaseRx.ReplaceAllStringFunc(s, func(match string) string {
		return strings.ToUpper(match[1:])
	})
}

func FieldName(field *protogen.Field) string {
	return Uncapitalize(SnakeToCamel(string(field.Desc.Name())))
}

func FieldNameType(field *protogen.Field) string {
	return string(field.Desc.Message().FullName())
}
