package java

import (
	"fmt"
	types "github.com/mtps/protoc-gen-cryptids/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
)

func BuilderSetEBytes(f *protogen.Field) string {
	fn := types.FieldName(f)
	setterName := fmt.Sprintf("set%s", types.Capitalize(fn))

	var c strings.Builder
	c.WriteString(fmt.Sprintf("/**\n"))
	c.WriteString(fmt.Sprintf(" * <code>.%s %s = %d;</code>\n", types.FieldNameType(f), f.Desc.Name(), f.Desc.Number()))
	c.WriteString(fmt.Sprintf(" */\n"))
	c.WriteString(fmt.Sprintf("public Builder %s(byte[] value, %s) {\n", setterName, EProviderParam))
	c.WriteString(fmt.Sprintf("  return %s(\n", setterName))
	c.WriteString(fmt.Sprintf("    %s.wrap(value, encryptionProvider)\n", types.PackageFor(types.TypeEBytes)))
	c.WriteString(fmt.Sprintf("  );\n"))
	c.WriteString(fmt.Sprintf("}\n"))
	c.WriteString(fmt.Sprintf("\n"))

	c.WriteString(fmt.Sprintf("/**\n"))
	c.WriteString(fmt.Sprintf(" * <code>.%s %s = %d;</code>\n", types.FieldNameType(f), f.Desc.Name(), f.Desc.Number()))
	c.WriteString(fmt.Sprintf(" */\n"))
	c.WriteString(fmt.Sprintf("public Builder %s(byte[] value) {\n", setterName))
	c.WriteString(fmt.Sprintf("  return %s(\n", setterName))
	c.WriteString(fmt.Sprintf("    %s.wrap(value)\n", types.PackageFor(types.TypeEBytes)))
	c.WriteString(fmt.Sprintf("  );\n"))
	c.WriteString(fmt.Sprintf("}\n"))
	c.WriteString(fmt.Sprintf("\n"))

	return c.String()
}

func EBytesDecrypt(retType string) string {
	pkg := types.PackageFor(types.TypeEBytes)
	c := strings.Builder{}

	// Instance methods
	// --

	// convert the current contained encrypted value into the decrypted version.
	c.WriteString("private byte[] decryptBytes(java.util.function.Function<byte[], byte[]> decryptionProvider) {\n")
	c.WriteString("  return decryptionProvider.apply(value_.toByteArray());\n")
	c.WriteString("}\n")
	c.WriteString("\n")

	// convert the current encrypted value into the decrypted version, using the default provider.
	c.WriteString("private byte[] decryptBytes() {\n")
	c.WriteString("  return decryptBytes(com.github.mtps.protobuf.crypt.CryptProvider.dec);\n")
	c.WriteString("}\n")
	c.WriteString("\n")

	// convert the current encrypted object type into the decrypted native type.
	c.WriteString(fmt.Sprintf("public byte[] decrypt(java.util.function.Function<byte[], byte[]> decryptionProvider) {\n"))
	c.WriteString("  return unwrap(decryptionProvider);\n")
	c.WriteString("}\n")
	c.WriteString("\n")

	c.WriteString(fmt.Sprintf("public byte[] decrypt() {\n"))
	c.WriteString("  return unwrap(com.github.mtps.protobuf.crypt.CryptProvider.dec);\n")
	c.WriteString("}\n")
	c.WriteString("\n")

	// --
	// Static methods.
	// --

	// wrap the decrypted value into a returned type of this EType.
	c.WriteString(fmt.Sprintf("public static %s wrap(byte[] value, java.util.function.Function<byte[], byte[]> encryptionProvider) {\n", pkg))
	c.WriteString(fmt.Sprintf("  return %s.newBuilder().setValue(\n", pkg))
	c.WriteString(fmt.Sprintf("    com.google.protobuf.ByteString.copyFrom(encryptionProvider.apply(value))\n"))
	c.WriteString(fmt.Sprintf("  ).build();\n"))
	c.WriteString(fmt.Sprintf("}\n"))
	c.WriteString("\n")

	// wrap the decrypted value into a returned type of this EType, using the default provider.
	c.WriteString(fmt.Sprintf("public static %s wrap(byte[] value) {\n", pkg))
	c.WriteString(fmt.Sprintf("  return wrap(value, com.github.mtps.protobuf.crypt.CryptProvider.enc);\n"))
	c.WriteString(fmt.Sprintf("}\n"))
	c.WriteString("\n")

	// unwrap the encrypted EType value into a returned native type.
	c.WriteString(fmt.Sprintf("private byte[] unwrap(java.util.function.Function<byte[], byte[]> decryptionProvider) {\n"))
	c.WriteString(fmt.Sprintf("  return decryptionProvider.apply(value_.toByteArray());\n"))
	c.WriteString(fmt.Sprintf("}\n"))
	c.WriteString("\n")

	// unwrap the encrypted EType value into a returned native type, using the default provider.
	c.WriteString(fmt.Sprintf("private byte[] unwrap() {\n"))
	c.WriteString(fmt.Sprintf("  return unwrap(com.github.mtps.protobuf.crypt.CryptProvider.enc);\n"))
	c.WriteString(fmt.Sprintf("}\n"))
	c.WriteString("\n")

	return c.String()
}
