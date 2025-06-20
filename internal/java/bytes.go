package java

import (
	"fmt"
	types "github.com/mtps/protoc-gen-cryptids/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
)

func BuilderSetEBytes(field *protogen.Field) string {
	return BuilderSet(field, "byte[]", types.PackageFor(types.TypeEBytes))
}

func BuilderSet(f *protogen.Field, retType string, pkg string) string {
	fn := types.FieldName(f)
	setterName := fmt.Sprintf("set%s", types.Capitalize(fn))

	var c strings.Builder
	c.WriteString(fmt.Sprintf("/**\n"))
	c.WriteString(fmt.Sprintf(" * <code>.%s %s = %d;</code>\n", types.FieldNameType(f), f.Desc.Name(), f.Desc.Number()))
	c.WriteString(fmt.Sprintf(" */\n"))
	c.WriteString(fmt.Sprintf("public Builder %s(%s value, %s) {\n", setterName, retType, EProviderParam))
	c.WriteString(fmt.Sprintf("  return %s(\n", setterName))
	c.WriteString(fmt.Sprintf("    %s.encrypt(value, encryptionProvider)\n", pkg))
	c.WriteString(fmt.Sprintf("  );\n"))
	c.WriteString(fmt.Sprintf("}\n"))
	c.WriteString(fmt.Sprintf("\n"))

	c.WriteString(fmt.Sprintf("/**\n"))
	c.WriteString(fmt.Sprintf(" * <code>.%s %s = %d;</code>\n", types.FieldNameType(f), f.Desc.Name(), f.Desc.Number()))
	c.WriteString(fmt.Sprintf(" */\n"))
	c.WriteString(fmt.Sprintf("public Builder %s(%s value) {\n", setterName, retType))
	c.WriteString(fmt.Sprintf("  return %s(\n", setterName))
	c.WriteString(fmt.Sprintf("    %s.encrypt(value)\n", pkg))
	c.WriteString(fmt.Sprintf("  );\n"))
	c.WriteString(fmt.Sprintf("}\n"))
	c.WriteString(fmt.Sprintf("\n"))
	return c.String()
}

func EBytesDecrypt() string {
	return ETypeDecrypt(
		"byte[]",
		"getValue().toByteArray()",
		"value",
	)
}

func EIntDecrypt() string {
	return ETypeDecrypt(
		"int",
		"java.nio.ByteBuffer.allocate(Integer.BYTES).putInt(value).order(java.nio.ByteOrder.BIG_ENDIAN).array()",
		"java.nio.ByteBuffer.wrap(value).getInt()",
	)
}

func ETypeDecrypt(retType string, toBytes string, fromBytes string) string {
	pkg := types.PackageFor(types.TypeEBytes)
	c := strings.Builder{}

	// --
	// Static methods.
	// --

	// Core type: fromBytes
	c.WriteString(fmt.Sprintf("public static %s fromBytes(byte[] value) {\n", retType))
	c.WriteString(fmt.Sprintf("  return %s;\n", fromBytes))
	c.WriteString("}\n")

	// Core type: toBytes
	c.WriteString(fmt.Sprintf("private byte[] toBytes() {\n"))
	c.WriteString(fmt.Sprintf("  return %s;\n", toBytes))
	c.WriteString("}\n")

	// wrap(value) -> EType with a provider.
	c.WriteString(fmt.Sprintf("public static %s encrypt(%s value, java.util.function.Function<byte[], byte[]> encryptionProvider) {\n", pkg, retType))
	c.WriteString(fmt.Sprintf("  return %s.newBuilder().setValue(\n", pkg))
	c.WriteString(fmt.Sprintf("    com.google.protobuf.ByteString.copyFrom(encryptionProvider.apply(value))\n"))
	c.WriteString(fmt.Sprintf("  ).build();\n"))
	c.WriteString(fmt.Sprintf("}\n"))
	c.WriteString("\n")

	// wrap(value) -> EType.
	c.WriteString(fmt.Sprintf("public static %s encrypt(%s value) {\n", pkg, retType))
	c.WriteString(fmt.Sprintf("  return encrypt(value, com.github.mtps.protobuf.crypt.CryptProvider.enc);\n"))
	c.WriteString(fmt.Sprintf("}\n"))
	c.WriteString("\n")

	// EType.unwrap() -> primitive with a provider
	c.WriteString(fmt.Sprintf("private %s decrypt(java.util.function.Function<byte[], byte[]> decryptionProvider) {\n", retType))
	c.WriteString(fmt.Sprintf("  return fromBytes(decryptionProvider.apply(getValue().toByteArray()));\n"))
	c.WriteString(fmt.Sprintf("}\n"))
	c.WriteString("\n")

	// EType.unwrap() -> primitive.
	c.WriteString(fmt.Sprintf("private %s decrypt() {\n", retType))
	c.WriteString(fmt.Sprintf("  return decrypt(com.github.mtps.protobuf.crypt.CryptProvider.enc);\n"))
	c.WriteString(fmt.Sprintf("}\n"))
	c.WriteString("\n")
	//
	//// Instance methods
	//// --
	//
	//// convert the current encrypted object type into the decrypted native type.
	//c.WriteString(fmt.Sprintf("public %s decrypt(java.util.function.Function<byte[], byte[]> decryptionProvider) {\n", retType))
	//c.WriteString("  return unwrap(decryptionProvider);\n")
	//c.WriteString("}\n")
	//c.WriteString("\n")
	//
	//c.WriteString(fmt.Sprintf("public %s decrypt() {\n", retType))
	//c.WriteString("  return unwrap(com.github.mtps.protobuf.crypt.CryptProvider.dec);\n")
	//c.WriteString("}\n")
	//c.WriteString("\n")
	//
	return c.String()
}
