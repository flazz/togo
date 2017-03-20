package main

import (
	"bytes"
	"flag"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

const chunkSize = 0x10

const tmpl = `
package {{.Pkg}}

var {{.Name}} = []byte{
    // {{ len .Value }} bytes from {{ .InputPath }}
	{{ range .Chunks -}}
		{{ range . }} {{ printf "0x%02x" . }}, {{ end }}
	{{ end }}
}
`

var t *template.Template

func init() {
	t = template.Must(template.New("constfile").Parse(tmpl))
}

type file struct {
	Pkg, Name, InputPath string
	Value                []byte
}

func (f *file) Chunks() [][]byte {
	return chunks(f.Value, chunkSize)
}

func chunks(b []byte, n int) [][]byte {
	var c [][]byte

	nChks := len(b) / n

	for i := 0; i < nChks; i++ {
		m := i * n
		c = append(c, b[m:m+n])
	}

	if r := len(b) % n; r > 0 {
		m := n * nChks
		c = append(c, b[m:m+r])
	}

	return c
}

func (f *file) Read() (err error) {
	f.Value, err = ioutil.ReadFile(f.InputPath)
	return
}

func (f *file) Render(w io.Writer) error {
	outputPath := f.InputPath + ".go"
	var buf bytes.Buffer
	if err := t.Execute(&buf, &f); err != nil {
		return err
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(outputPath, b, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func main() {

	var f file
	flag.StringVar(&f.Pkg, "pkg", "", "package")
	flag.StringVar(&f.Name, "name", "", "const name")
	flag.StringVar(&f.InputPath, "input", "", "input file")
	flag.Parse()

	if f.Pkg == "" {
		log.Fatal("pkg required")
	}

	if f.Name == "" {
		log.Fatal("name required")
	}

	if f.InputPath == "" {
		log.Fatal("input file required")
	}

	if err := f.Read(); err != nil {
		log.Fatal(err)
	}

	if err := f.Render(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
