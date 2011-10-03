package main

import (
	"flag"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

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
		pDir := path.Join(dir, "out", name)
		tDir := path.Join(dir, "themes", theme)
		CopyDir(path.Join(pDir, "theme"), tDir)

		os.MkdirAll(pDir, 0755)
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
