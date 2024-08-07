package main

import (
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/sprig/v3"
)

var tmpl *template.Template

func parseTemplates() error {
	t := template.New("").Funcs(sprig.FuncMap())
	err := filepath.Walk("templates", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(path, ".html") {
			tmplBytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			_, err = t.New(filepath.Base(path)).Parse(string(tmplBytes))
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	tmpl = t
	return nil
}
