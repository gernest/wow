// +build mage

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/magefile/mage/mg"
)

func Spinners() error {
	mg.Deps(func() error {
		pkg := "spinner"
		_, err := os.Stat(pkg)
		if err != nil {
			if os.IsNotExist(err) {
				return os.Mkdir(pkg, 0777)
			}
			return err
		}
		return nil
	})
	b, err := ioutil.ReadFile("spinners.json")
	if err != nil {
		return err
	}
	o := make(map[string]interface{})
	err = json.Unmarshal(b, &o)
	if err != nil {
		return err
	}
	tpl, err := template.New("spinner").Funcs(helpers()).Parse(spinnersTpl)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, o)
	if err != nil {
		return err
	}
	bo, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	return ioutil.WriteFile("spinner/spinners.go", bo, 0600)
}

func helpers() template.FuncMap {
	return template.FuncMap{
		"title": strings.Title,
		"stringify": func(a []interface{}) string {
			o := ""
			switch len(a) {
			case 0:
				return ""
			case 1:
				return fmt.Sprintf("\"%s\"", a[0])
			default:
				for k, v := range a {
					if k == 0 {
						o += fmt.Sprintf("`%s`", v)
					} else {
						if v == "`" {
							o += fmt.Sprintf(",\"%s\"", v)

						} else {
							o += fmt.Sprintf(",`%s`", v)
						}

					}
				}

				return fmt.Sprintf("[]string{ %v }", o)
			}
		},
	}
}

const spinnersTpl = `
package spinner

// Spinner defines a spinner widget
type Spinner struct{
	Name Name 
	Interval int
	Frames []string
}

// Name  represents a name for a spinner item.
type Name uint

// available spinners
const(
	Unknown Name=iota
	{{range $k,$v:= .}}
	{{- $k|title}}
	{{end}}
)

func (s Name)String()string{
	switch s{
		{{- range $k,$v:=.}}
	case {{$k|title}} :
		return "{{$k}}"
		{{- end}}
	default:
		return ""
	}
}

func GetSpinner( name Name)Spinner{
	switch name{
		{{- range $k,$v:=.}}
	case {{$k|title}} :
		return Spinner{
			Name: {{$k|title}},
			Interval: {{$v.interval}},
			Frames: {{$v.frames|stringify }},
		}
		{{- end}}
	default:
		return Spinner{}
	}
}
`
