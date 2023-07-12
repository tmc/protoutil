package main

import (
	"flag"
	"io"
	"log"
	"os"
	"text/template"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

var flagTemplate = flag.String("template", "", "template expression")

func main() {
	flag.Parse()
	r := &pluginpb.CodeGeneratorRequest{}

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if err := proto.Unmarshal(data, r); err != nil {
		log.Fatal(err)
	}
	if *flagTemplate == "" {
		b, err := protojson.Marshal(r)
		if err != nil {
			log.Fatal(err)
		}
		os.Stderr.Write(b)
	} else {
		if err := render(os.Stderr, *flagTemplate, r); err != nil {
			log.Fatal(err)
		}
	}

	os.Stdout.Write(data)
}

func render(w io.Writer, templateStr string, r *pluginpb.CodeGeneratorRequest) error {
	t := template.Must(template.New("").Parse(templateStr))
	return t.Execute(w, r)
}
