package main

import (
	"fmt"
	"github.com/mtps/protoc-gen-cryptids/internal"
	"github.com/mtps/protoc-gen-cryptids/internal/java"
	"github.com/mtps/protoc-gen-cryptids/internal/kotlin"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/pluginpb"
	"io"
	"log"
	"os"
	"strings"
)

func inputRequest(opts protogen.Options) (*protogen.Plugin, error) {
	if len(os.Args) > 1 {
		return nil, fmt.Errorf("unknown argument %q (this program should be run by protoc, not directly)", os.Args[1])
	}
	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	req := &pluginpb.CodeGeneratorRequest{}
	if err := proto.Unmarshal(in, req); err != nil {
		return nil, err
	}
	gen, err := opts.New(req)
	if err != nil {
		return nil, err
	}
	return gen, nil
}

func outputResponse(resp *pluginpb.CodeGeneratorResponse) error {
	out, err := proto.Marshal(resp)
	if err != nil {
		return err
	}
	if _, err := os.Stdout.Write(out); err != nil {
		return err
	}
	return nil
}

func main() {
	opts := protogen.Options{}

	gen, err := inputRequest(opts)
	if err != nil {
		log.Fatal(err)
	}

	paramKv := make(map[string]string)
	params := strings.Split(gen.Request.GetParameter(), ":")
	for _, p := range params {
		k, v, _ := strings.Cut(p, "=")
		fmt.Fprintf(os.Stderr, "%s=%s\n", k, v)
		paramKv[k] = v
	}

	// The type needs to be split since we can only output kotlin code into the
	// src/main/kotlin folder if defined on the command line via output param.
	//
	// Base java class modification by default.
	genJava := true
	// Kotlin extension off by default.
	genKt := false
	if paramKv["g"] == "java" {
		genJava = true
		genKt = false
	} else if paramKv["g"] == "kotlin" {
		genKt = true
		genJava = false
	}

	// Emit the driver for providers to hook into.
	if genJava {
		java.GenCryptProvider(gen)
	}

	var addlFiles []*pluginpb.CodeGeneratorResponse_File
	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}

		if genKt {
			generateKotlinFile(gen, f)
		}

		fmt.Fprintf(os.Stderr, "processing %s\n", f.Desc.Path())
		for _, m := range f.Messages {
			if m.Desc.FullName() == types.TypeEBytes.FullName() || m.Desc.FullName() == types.TypeEInt.FullName() {
				if genJava {
					var fun func() string
					switch m.Desc.FullName() {
					case types.TypeEBytes.FullName():
						fun = java.EBytesDecrypt
						break
					case types.TypeEInt.FullName():
						fun = java.EIntDecrypt
						break
					}

					gf := generateJavaCryptHelpers(gen, f, m, fun)
					fmt.Fprintf(os.Stderr, "file:%s ip:%s :%s\n", gf.GetName(), gf.GetInsertionPoint(), gf.GetContent())
					addlFiles = append(addlFiles, gf)
				}
				continue
			}

			if m.Desc.ParentFile().Package() == "crypt" {
				fmt.Fprintf(os.Stderr, "skipping internal type:%s\n", m.Desc.Name())
				continue
			}

			fmt.Fprintf(os.Stderr, " - m: %s\n", m.Desc.Name())

			if genJava {
				gf := generateJavaBuilderCryptSetters(gen, f, m)
				fmt.Fprintf(os.Stderr, "file:%s ip:%s :%s\n", gf.GetName(), gf.GetInsertionPoint(), gf.GetContent())
				addlFiles = append(addlFiles, gf)
			}

			if genKt {

			}
		}
	}

	resp := gen.Response()
	resp.File = append(resp.File, addlFiles...)
	if err := outputResponse(resp); err != nil {
		log.Fatal(err)
	}
}

func generateJavaCryptHelpers(gen *protogen.Plugin, f *protogen.File, m *protogen.Message, ccc func() string) *pluginpb.CodeGeneratorResponse_File {
	fa, _ := strings.CutSuffix(string(f.Desc.Name()), ".")
	javaPkgBase := strings.ReplaceAll(*f.Proto.Options.JavaPackage, ".", "/")
	fileName := javaPkgBase + "/" + types.Capitalize(fa) + "Proto.java"
	iPoint := "class_scope:" + string(m.Desc.FullName())

	var c strings.Builder
	c.WriteString(ccc())

	content := c.String()
	return &pluginpb.CodeGeneratorResponse_File{
		Name:           &fileName,
		InsertionPoint: &iPoint,
		Content:        &content,
	}
}

func generateJavaBuilderCryptSetters(gen *protogen.Plugin, f *protogen.File, m *protogen.Message) *pluginpb.CodeGeneratorResponse_File {
	fa, _ := strings.CutSuffix(string(f.Desc.Name()), ".")
	javaPkgBase := strings.ReplaceAll(*f.Proto.Options.JavaPackage, ".", "/")
	fileName := javaPkgBase + "/" + types.Capitalize(fa) + ".java"
	iPoint := "builder_scope:" + string(m.Desc.FullName())

	var content strings.Builder
	for _, field := range m.Fields {
		if field.Desc.Kind() != protoreflect.MessageKind {
			continue
		}

		fieldName := types.FieldName(field)
		fieldTypeName := types.FieldNameType(field)
		switch fieldTypeName {
		case string(types.TypeEString.FullName()):
			fmt.Fprintf(os.Stderr, "detected encrypted string field %s\n", fieldName)

			break

		case string(types.TypeEBytes.FullName()):
			fmt.Fprintf(os.Stderr, "detected encrypted bytes field %s\n", fieldName)
			content.WriteString(java.BuilderSetEBytes(field))
			break

		case string(types.TypeETimestamp.FullName()):
			fmt.Fprintf(os.Stderr, "detected encrypted timestamp field %s\n", fieldName)

			break

		case string(types.TypeEInt.FullName()):
			fmt.Fprintf(os.Stderr, "detected encrypted int field %s\n", fieldName)

			break

		case string(types.TypeEAny.FullName()):
			fmt.Fprintf(os.Stderr, "detected encrypted anything field %s\n", fieldName)

			break
		}
	}

	c := content.String()
	return &pluginpb.CodeGeneratorResponse_File{
		Name:           &fileName,
		InsertionPoint: &iPoint,
		Content:        &c,
	}
}

func generateKotlinFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	fmt.Fprintf(os.Stderr, "generating for %s\n", file.Desc.Path())

	var cryptJavaClass string
	if file.Proto.Options.JavaOuterClassname != nil {
		cryptJavaClass = *file.Proto.Options.JavaOuterClassname
	} else {
		cryptJavaClass = types.Capitalize(string(file.Desc.Name()))
	}

	javaPackage := *file.Proto.Options.JavaPackage
	javaDirectory := strings.ReplaceAll(javaPackage, ".", "/")
	filename := javaDirectory + "/" + cryptJavaClass + "CryptKt.kt"

	fmt.Fprintf(os.Stderr, "javaDir:%s msgClass:%s javaPkg:%s filename:%s  goimportpath:%s\n", javaDirectory, cryptJavaClass, javaPackage, filename, file.GoImportPath)
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-cryptids")
	g.P("// source: ", file.Desc.Path())
	g.P()
	g.P("package ", javaPackage)
	g.P()
	for _, message := range file.Messages {
		typeName := cryptJavaClass + "." + string(message.Desc.Name())
		fmt.Fprintf(os.Stderr, "typeName:%s\n", typeName)
		for _, field := range message.Fields {
			// Handle EStrings
			if field.Desc.Kind() == protoreflect.MessageKind {
				fieldTypeName := string(field.Desc.Message().FullName())
				fieldName := types.Capitalize(types.SnakeToCamel(string(field.Desc.Name())))
				switch fieldTypeName {
				case string(types.TypeEString.FullName()):
					fmt.Fprintf(os.Stderr, "detected encrypted string field %s\n", fieldName)
					kotlin.EmitEStringWrapper(g, fieldName, typeName)
					break

				case string(types.TypeEBytes.FullName()):
					fmt.Fprintf(os.Stderr, "detected encrypted bytes field %s\n", fieldName)
					kotlin.EmitEBytesWrapper(g, fieldName, typeName)
					break

				case string(types.TypeETimestamp.FullName()):
					fmt.Fprintf(os.Stderr, "detected encrypted timestamp field %s\n", fieldName)
					kotlin.EmitETimestampWrapper(g, fieldName, typeName)
					break

				case string(types.TypeEInt.FullName()):
					fmt.Fprintf(os.Stderr, "detected encrypted int field %s\n", fieldName)
					kotlin.EmitEIntWrapper(g, fieldName, typeName)
					break

				case string(types.TypeEAny.FullName()):
					fmt.Fprintf(os.Stderr, "detected encrypted anything field %s\n", fieldName)
					kotlin.EmitEAnyWrapper(g, fieldName, typeName)
					break
				}
			}
		}
	}
	fmt.Fprintf(os.Stderr, "writing file %s\n", filename)
	return g
}
