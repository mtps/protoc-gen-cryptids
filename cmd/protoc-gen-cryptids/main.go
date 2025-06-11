package main

import (
	// "fmt"
	_ "github.com/FigureTechnologies/protobuf/protoc-gen-cryptids/crypt"
	// "go/format"
	// plugin "google.golang.org/protobuf/compiler/protogen"
	// "io/ioutil"
	// "log"
	// "os"
	// "path/filepath"
	// "reflect"
	// "strings"
)

//	func Read() *plugin.CodeGeneratorRequest {
//		g := generator.New()
//		data, err := ioutil.ReadAll(os.Stdin)
//		if err != nil {
//			g.Error(err, "reading input")
//		}
//
//		if err := proto.Unmarshal(data, g.Request); err != nil {
//			g.Error(err, "parsing input proto")
//		}
//
//		if len(g.Request.FileToGenerate) == 0 {
//			g.Fail("no files to generate")
//		}
//		return g.Request
//	}
//
//	func goformat(resp *plugin.CodeGeneratorResponse) error {
//		for i := 0; i < len(resp.File); i++ {
//			formatted, err := format.Source([]byte(resp.File[i].GetContent()))
//			if err != nil {
//				return fmt.Errorf("go format error: %v", err)
//			}
//			fmts := string(formatted)
//			resp.File[i].Content = &fmts
//		}
//		return nil
//	}
//
// func getProto(messageType string, messageBytes []byte) proto.Message {
//
//		pbtype := proto.MessageType(messageType)
//		if pbtype == nil {
//			panic(messageType)
//		}
//		msg := reflect.New(pbtype.Elem()).Interface().(proto.Message)
//		proto.Unmarshal(messageBytes, msg)
//		return msg
//	}
//
//	var flog = func(format string, a ...interface{}) {
//		_, _ = fmt.Fprintf(os.Stderr, format, a...)
//	}
//
//	func Generate(req *plugin.CodeGeneratorRequest) (res plugin.CodeGeneratorResponse) {
//		for _, f := range req.ProtoFile {
//			if f.Package == nil {
//				p := ""
//				f.Package = &p
//			}
//
//			newName := fmt.Sprintf("%s.properties", strings.TrimSuffix(*f.Name, filepath.Ext(*f.Name)))
//			var lines []string
//			for _, m := range f.MessageType {
//				if m.Options == nil {
//					// No options, disregard.
//					continue
//				}
//
//				fqn := *f.Package + "." + *m.Name
//				opts := m.GetOptions()
//				if opts == nil || !proto.HasExtension(opts, kafkapb.E_Binding) {
//					continue
//				}
//
//				flog("\n")
//				flog("# %s (%s)\n", fqn, *f.Name)
//				ext, err := proto.GetExtension(opts, kafkapb.E_Binding)
//				if err != nil {
//					panic(err)
//				}
//
//				topics, ok := ext.([]*kafkapb.TopicBinding)
//				if !ok {
//					panic("topics not []*kafkapb.TopicBinding type!")
//				}
//
//				for _, topic := range topics {
//					comment := fmt.Sprintf("# t:%s id:'%s'", fqn, topic.Id)
//					content := fmt.Sprintf("%s=%s", topic.Topic, fqn)
//					lines = append(lines, comment)
//					lines = append(lines, content)
//					flog("# ~> %s\n", topic)
//				}
//				flog("\n")
//			}
//
//			if len(lines) > 0 {
//				content := strings.Join(lines, "\n") + "\n"
//				res.File = append(res.File, &plugin.CodeGeneratorResponse_File{
//					Name:    &newName,
//					Content: &content,
//				})
//			}
//		}
//
//		return res
//	}
//
// // Error reports a problem, including an error, and exits the program.
//
//	func Error(err error, msgs ...string) {
//		s := strings.Join(msgs, " ") + ":" + err.Error()
//		log.Print("protoc-gen-kafka: error:", s)
//		os.Exit(1)
//	}
//
//	func Write(resp plugin.CodeGeneratorResponse) {
//		// Send back the results.
//		data, err := proto.Marshal(&resp)
//		if err != nil {
//			Error(err, "failed to marshal output proto")
//		}
//		_, err = os.Stdout.Write(data)
//		if err != nil {
//			Error(err, "failed to write output proto")
//		}
//	}
func main() {
	//	Write(Generate(Read()))
}
