package kotlin

import (
	types "github.com/mtps/protoc-gen-cryptids/internal"
	"google.golang.org/protobuf/compiler/protogen"
)

func EmitEStringWrapper(g *protogen.GeneratedFile, fieldName string, typeName string) {
	g.P("fun ", typeName, ".decrypt", types.Capitalize(fieldName), "(decryptionProvider: (ByteArray) -> ByteArray = ", types.CryptProviderName, ".dec): String =")
	g.P("    ", StringFromEString(fieldName))
	g.P()
	g.P("fun ", typeName, ".Builder.set", types.Capitalize(fieldName), "(value: String, encryptionProvider: (ByteArray) -> ByteArray = ", types.CryptProviderName, ".enc): ", typeName, ".Builder =")
	g.P("    set", types.Capitalize(fieldName), "(", EStringFromString("value"), ")")
	g.P()
}

func EmitEBytesWrapper(g *protogen.GeneratedFile, fieldName string, typeName string) {
	g.P("fun ", typeName, ".decrypt", types.Capitalize(fieldName), "(decryptionProvider: (ByteArray) -> ByteArray = ", types.CryptProviderName, ".dec): ByteArray =")
	g.P("    ", BytesFromEBytes(types.Uncapitalize(fieldName)))
	g.P()
	g.P("fun ", typeName, ".Builder.set", fieldName, "(value: ByteArray, encryptionProvider: (ByteArray) -> ByteArray = ", types.CryptProviderName, ".enc): ", typeName, ".Builder =")
	g.P("    set", types.Capitalize(fieldName), "(", EBytesFromBytes("value"), ")")
	g.P()
}

func EmitETimestampWrapper(g *protogen.GeneratedFile, fieldName string, typeName string) {
	g.P("fun ", typeName, ".decrypt", types.Capitalize(fieldName), "(decryptionProvider: (ByteArray) -> ByteArray = ", types.CryptProviderName, ".dec): com.google.protobuf.Timestamp =")
	g.P("    ", TimestampFromETimestamp(types.Uncapitalize(fieldName)))
	g.P()
	g.P("fun ", typeName, ".Builder.set", types.Capitalize(fieldName), "(value: com.google.protobuf.Timestamp, encryptionProvider: (ByteArray) -> ByteArray = ", types.CryptProviderName, ".enc): ", typeName, ".Builder =")
	g.P("    set", types.Capitalize(fieldName), "(", ETimestampFromTimestamp("value"), ")")
	g.P()
}

func EmitEIntWrapper(g *protogen.GeneratedFile, fieldName string, typeName string) {
	g.P("fun ", typeName, ".decrypt", fieldName, "(decryptionProvider: (ByteArray) -> ByteArray = ", types.CryptProviderName, ".dec): Int =")
	g.P("    ", IntFromEInt(types.Uncapitalize(fieldName)))
	g.P()
	g.P("fun ", typeName, ".Builder.set", fieldName, "(value: Int, encryptionProvider: (ByteArray) -> ByteArray = ", types.CryptProviderName, ".enc): ", typeName, ".Builder =")
	g.P("    set", types.Capitalize(fieldName), "(", EIntFromInt("value"), ")")
	g.P()
}

func EmitEAnyWrapper(g *protogen.GeneratedFile, fieldName string, typeName string) {
	g.P("fun ", typeName, ".decrypt", fieldName, "(decryptionProvider: (ByteArray) -> ByteArray = ", types.CryptProviderName, ".dec): com.google.protobuf.Any = ")
	g.P("    ", AnyFromEAny(types.Uncapitalize(fieldName)))
	g.P()
	g.P("fun ", typeName, ".Builder.set", fieldName, "(value: com.google.protobuf.Any, encryptionProvider: (ByteArray) -> ByteArray = ", types.CryptProviderName, ".enc): ", typeName, ".Builder =")
	g.P("    set", types.Capitalize(fieldName), "(", EAnyFromAny("value"), ")")
	g.P()
}
