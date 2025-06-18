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

	// Emit the driver for providers to hook into.
	java.GenCryptProvider(gen)

	var addlFiles []*pluginpb.CodeGeneratorResponse_File
	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}

		fmt.Fprintf(os.Stderr, "processing %s\n", f.Desc.Path())

		fa, _ := strings.CutSuffix(string(f.Desc.Name()), ".")
		for _, m := range f.Messages {
			if m.Desc.FullName() == types.TypeEBytes.FullName() {
				gf := generateJavaCryptHelpers(gen, f, m)
				addlFiles = append(addlFiles, gf)
				continue
			}

			if m.Desc.ParentFile().Package() == "crypt" {
				fmt.Fprintf(os.Stderr, "skipping internal type:%s\n", m.Desc.Name())
				continue
			}

			fmt.Fprintf(os.Stderr, " - m: %s\n", m.Desc.Name())

			gf := generateJavaBuilderCryptSetters(gen, f, m)

			fmt.Fprintf(os.Stderr, "fa:%s fn:%s ip:%s\n", fa, gf.GetName(), gf.GetInsertionPoint())
			addlFiles = append(addlFiles, gf)
		}
	}
	resp := gen.Response()
	resp.File = append(resp.File, addlFiles...)
	if err := outputResponse(resp); err != nil {
		log.Fatal(err)
	}
}

func generateJavaCryptHelpers(gen *protogen.Plugin, f *protogen.File, m *protogen.Message) *pluginpb.CodeGeneratorResponse_File {
	fa, _ := strings.CutSuffix(string(f.Desc.Name()), ".")
	javaPkgBase := strings.ReplaceAll(*f.Proto.Options.JavaPackage, ".", "/")
	fileName := javaPkgBase + "/" + types.Capitalize(fa) + "Proto.java"
	iPoint := "class_scope:" + string(m.Desc.FullName())

	typeBytes := types.PackageFor(types.TypeEBytes)

	var c strings.Builder
	c.WriteString(java.EBytesDecrypt(typeBytes))

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

func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
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
