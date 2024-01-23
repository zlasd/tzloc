package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var (
	OutputFile string
)

type LineParam struct {
	Identifier string
	Value      string
}

func main() {
	flag.StringVar(&OutputFile, "output", "", "output file path")
	flag.Parse()

	tzDataPath := GetTzDataPath()
	tzList := LoadTzList(tzDataPath)
	lines := make([]LineParam, 0, len(tzList))
	for _, tz := range tzList {
		//fmt.Println(GetIdentifierName(tz), "=", tz)
		lines = append(lines, LineParam{
			Identifier: GetIdentifierName(tz),
			Value:      tz,
		})
	}

	tmplStr := `package tzloc

const ({{range .}}
	{{.Identifier}} = "{{.Value}}"{{end}}
)

var locationMap = map[string]struct{}{
{{range .}}	{{.Identifier}}: {},
{{end}}}
`
	var buf []byte
	buffer := bytes.NewBuffer(buf)
	tmpl := template.Must(template.New("timezone").Parse(tmplStr))
	err := tmpl.Execute(buffer, lines)
	if err != nil {
		log.Fatalln(err)
	}
	out := buffer.Bytes()
	if err := os.WriteFile(OutputFile, out, 0644); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(out))
}

func GetTzDataPath() string {
	out, err := exec.Command("go", "env", "GOROOT").Output()
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Trim(string(out), "\n") + "/lib/time/zoneinfo.zip"
}

func LoadTzList(path string) []string {
	fmt.Println("Exec zipinfo command: zipinfo -1", path)
	out, err := exec.Command("zipinfo", "-1", path).Output()
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Split(strings.Trim(string(out), "\n"), "\n")
}

func GetIdentifierName(tzVal string) string {
	replacer := strings.NewReplacer(
		"/", "_",
		"GMT-", "GMT_Minus_",
		"-", "_",
		"+", "_Plus_",
	)
	return replacer.Replace(tzVal)
}
