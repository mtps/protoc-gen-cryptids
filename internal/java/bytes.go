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
		types.PackageFor(types.TypeEBytes),
		"return value;",
		"return value;",
	)
}

func EIntDecrypt() string {
	return ETypeDecrypt(
		"int",
		types.PackageFor(types.TypeEInt),
		"return java.nio.ByteBuffer.allocate(Integer.BYTES).putInt(value).order(java.nio.ByteOrder.BIG_ENDIAN).array();",
		"return java.nio.ByteBuffer.wrap(value).getInt();",
	)
}

func EStringDecrypt() string {
	return ETypeDecrypt(
		"String",
		types.PackageFor(types.TypeEString),
		"return value.getBytes();",
		"return new String(value);",
	)
}

func ETimestampDecrypt() string {
	return ETypeDecrypt(
		"com.google.protobuf.Timestamp",
		types.PackageFor(types.TypeETimestamp),
		"return java.nio.ByteBuffer.allocate(Integer.BYTES + Long.BYTES).putLong(value.getSeconds()).putInt(value.getNanos()).array();",
		"java.nio.ByteBuffer bb = java.nio.ByteBuffer.wrap(value).order(java.nio.ByteOrder.BIG_ENDIAN);\n"+
			"  return com.google.protobuf.Timestamp.newBuilder().setSeconds(bb.getLong()).setNanos(bb.getInt()).build();",
	)
}

func EAnyDecrypt() string {
	return ETypeDecrypt(
		"com.google.protobuf.Any",
		types.PackageFor(types.TypeEAny),
		"return value.toByteArray();",
		"try { return com.google.protobuf.Any.parseFrom(value); } "+
			"catch (com.google.protobuf.InvalidProtocolBufferException e) { "+
			"throw new RuntimeException(\"Failed to convert from bytes to ProtoAny\", e); }",
	)
}

func ETypeDecrypt(retType string, eType string, toBytes string, fromBytes string) string {

	c := strings.Builder{}

	// --
	// Static methods.
	// --

	// Core type: fromBytes
	c.WriteString(fmt.Sprintf("public static %s fromBytes(byte[] value) {\n", retType))
	c.WriteString(fmt.Sprintf("  %s\n", fromBytes))
	c.WriteString("}\n\n")

	// Core type: toBytes
	c.WriteString(fmt.Sprintf("private static byte[] toBytes(%s value) {\n", retType))
	c.WriteString(fmt.Sprintf("  %s\n", toBytes))
	c.WriteString("}\n\n")

	// wrap(value) -> EType with a provider.
	c.WriteString(fmt.Sprintf("public static %s encrypt(%s value, java.util.function.Function<byte[], byte[]> encryptionProvider) {\n", eType, retType))
	c.WriteString(fmt.Sprintf("  if (encryptionProvider == null) {\n"))
	c.WriteString(fmt.Sprintf("    throw new NullPointerException();\n"))
	c.WriteString(fmt.Sprintf("  }\n"))
	c.WriteString(fmt.Sprintf("  return %s.newBuilder().setValue(\n", eType))
	c.WriteString(fmt.Sprintf("    com.google.protobuf.ByteString.copyFrom(encryptionProvider.apply(toBytes(value)))\n"))
	c.WriteString(fmt.Sprintf("  ).build();\n"))
	c.WriteString(fmt.Sprintf("}\n\n"))

	// wrap(value) -> EType.
	c.WriteString(fmt.Sprintf("public static %s encrypt(%s value) {\n", eType, retType))
	c.WriteString(fmt.Sprintf("  return encrypt(value, com.github.mtps.protobuf.crypt.CryptProviderRegistry.enc);\n"))
	c.WriteString(fmt.Sprintf("}\n\n"))

	// --
	// Instance methods
	// --

	// EType.unwrap() -> primitive with a provider
	c.WriteString(fmt.Sprintf("public %s decrypt(java.util.function.Function<byte[], byte[]> decryptionProvider) {\n", retType))
	c.WriteString(fmt.Sprintf("  if (decryptionProvider == null) {\n"))
	c.WriteString(fmt.Sprintf("    throw new NullPointerException();\n"))
	c.WriteString(fmt.Sprintf("  }\n"))
	c.WriteString(fmt.Sprintf("  return fromBytes(decryptionProvider.apply(getValue().toByteArray()));\n"))
	c.WriteString(fmt.Sprintf("}\n"))
	c.WriteString("\n")

	// EType.unwrap() -> primitive.
	c.WriteString(fmt.Sprintf("public %s decrypt() {\n", retType))
	c.WriteString(fmt.Sprintf("  return decrypt(com.github.mtps.protobuf.crypt.CryptProviderRegistry.enc);\n"))
	c.WriteString(fmt.Sprintf("}\n"))
	c.WriteString("\n")
	return c.String()
}
