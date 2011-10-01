package main

import (
	"flag"
	"os"
	"path"
)

func copyDir(dst, src string) {
	// FIXME: need to copy from a directory to another
}

func generatePresentation(name, theme string) {
	if dir, err := os.Getwd(); err == nil {
		pDir := path.Join(dir, "out", name)
		tDir := path.Join(dir, "themes", theme)
		copyDir(path.Join(pDir, "theme"), tDir)

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
		generatePresentation(*name, *theme)
	}
}
