package kotlin

import (
	types "github.com/mtps/protoc-gen-cryptids/internal"
	"google.golang.org/protobuf/compiler/protogen"
)

func EmitEStringWrapper(g *protogen.GeneratedFile, typeName string) {
	typ := types.PackageFor(types.TypeEString)
	g.P("public fun ", typeName, "Kt.Dsl.encrypt(")
	g.P("    value: String,")
	g.P("    encryptionProvider: ((ByteArray) -> ByteArray)? = null) =")
	g.P("  if (encryptionProvider == null) {")
	g.P("    ", typ, ".encrypt(value)")
	g.P("  } else {")
	g.P("    ", typ, ".encrypt(value, encryptionProvider)")
	g.P("  }")
	g.P()
}

func EmitEBytesWrapper(g *protogen.GeneratedFile, typeName string) {
	typ := types.PackageFor(types.TypeEBytes)
	g.P("public fun ", typeName, "Kt.Dsl.encrypt(")
	g.P("    value: ByteArray,")
	g.P("    encryptionProvider: ((ByteArray) -> ByteArray)? = null) =")
	g.P("  if (encryptionProvider == null) {")
	g.P("    ", typ, ".encrypt(value)")
	g.P("  } else {")
	g.P("    ", typ, ".encrypt(value, encryptionProvider)")
	g.P("  }")
	g.P()
}

func EmitETimestampWrapper(g *protogen.GeneratedFile, typeName string) {
	typ := types.PackageFor(types.TypeETimestamp)
	g.P("public fun ", typeName, "Kt.Dsl.encrypt(")
	g.P("    value: String,")
	g.P("    encryptionProvider: ((ByteArray) -> ByteArray)? = null) =")
	g.P("  if (encryptionProvider == null) {")
	g.P("    ", typ, ".encrypt(value)")
	g.P("  } else {")
	g.P("    ", typ, ".encrypt(value, encryptionProvider)")
	g.P("  }")
	g.P()
}

func EmitEIntWrapper(g *protogen.GeneratedFile, typeName string) {
	typ := types.PackageFor(types.TypeEInt)
	g.P("public fun ", typeName, "Kt.Dsl.encrypt(")
	g.P("    value: String,")
	g.P("    encryptionProvider: ((ByteArray) -> ByteArray)? = null) =")
	g.P("  if (encryptionProvider == null) {")
	g.P("    ", typ, ".encrypt(value)")
	g.P("  } else {")
	g.P("    ", typ, ".encrypt(value, encryptionProvider)")
	g.P("  }")
	g.P()
}

func EmitEAnyWrapper(g *protogen.GeneratedFile, typeName string) {
	typ := types.PackageFor(types.TypeEAny)
	g.P("public fun ", typeName, "Kt.Dsl.encrypt(")
	g.P("    value: String,")
	g.P("    encryptionProvider: ((ByteArray) -> ByteArray)? = null) =")
	g.P("  if (encryptionProvider == null) {")
	g.P("    ", typ, ".encrypt(value)")
	g.P("  } else {")
	g.P("    ", typ, ".encrypt(value, encryptionProvider)")
	g.P("  }")
	g.P()
}
