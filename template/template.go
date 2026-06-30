package template

import (
	"bytes"
	"io"
	"path/filepath"
	"text/template"
)

type Template struct {
	tpl   *template.Template
	funcs template.FuncMap
	wr    io.Writer
}

func Load(fileName string) (Template, error) {
	var err error
	name := filepath.Base(fileName)
	t := Template{funcs: funcs(), wr: writer(), tpl: template.New(name)}
	t.tpl, err = t.tpl.ParseFiles(fileName)
	if err != nil {
		return Template{}, err
	}
	return t, nil
}

func (t Template) Execute(name string, data any) () {
	buf := &bytes.Buffer{}
	t.tpl.ExecuteTemplate(buf, name, data)
}

func funcs() template.FuncMap {
	return nil
}

func writer() io.Writer {
	return nil
}