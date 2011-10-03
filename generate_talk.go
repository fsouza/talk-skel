package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"template"
)

type TemplateData struct {
	Name string
}

func CopyDir(dst, src string) {
	filepath.Walk(src, func (path string, info *os.FileInfo, err os.Error) (os.Error) {
		file := strings.Replace(path, src, dst, -1)

		if info.IsDirectory() {
			os.MkdirAll(file, 0755)
		} else {
			srcFile, errSrc := os.Open(path)
			defer srcFile.Close()
			if errSrc != nil {
				return errSrc
			}

			dstFile, errDst := os.Create(file)
			defer dstFile.Close()
			if errDst != nil {
				return errDst
			}

			io.Copy(dstFile, srcFile)
		}

		return nil
	})
}

func GeneratePresentation(name, theme string) {
	if dir, err := os.Getwd(); err == nil {
		out := path.Join(dir, "out", name)
		tDir := path.Join(dir, "themes", theme)
		CopyDir(out, path.Join(dir, "resources"))
		CopyDir(path.Join(out, "theme"), tDir)

		templates := map[string] string{
			"cfg.tpl": fmt.Sprintf("%s.cfg", name),
			"fabfile.py.tpl": "fabfile.py",
			"rst.tpl": fmt.Sprintf("%s.rst", name),
		}

		data := &TemplateData{name}
		templatesDir := path.Join(dir, "templates")
		for k, v := range templates {
			t, err := template.ParseFile(path.Join(templatesDir, k))

			if err != nil {
				panic(err)
			}

			dstFile, _ := os.Create(path.Join(out, v))
			defer dstFile.Close()

			t.Execute(dstFile, data)
		}
	}
}

func main() {
	name := flag.String("name", "", "The presentation name (required)")
	theme := flag.String("theme", "default", "The theme to be used (available in the themes directory)")
	flag.Parse()

	if *name == "" {
		flag.Usage()
	} else {
		GeneratePresentation(*name, *theme)
	}
}
