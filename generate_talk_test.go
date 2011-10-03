package main

import (
	"os"
	"path"
	"strings"
	"testing"
)

func TestCopyDirectoryShouldCreateTheDestinationDirectory(t *testing.T) {
	src := "/tmp/src"
	dst := "/tmp/dst"
	os.MkdirAll(src, 0755)
	CopyDir(dst, src)

	info, _ := os.Stat(dst)
	if !info.IsDirectory() {
		t.Errorf("CopyDir didn't created the destination directory, and it should!")
	}

	os.RemoveAll(src)
	os.RemoveAll(dst)
}

func TestCopyDirectoryShouldCopyTheDirectContent(t *testing.T) {
	src := "/tmp/src"
	complete := path.Join(src, "bla", "bla", "blabla")
	dst := "/tmp/dst"

	os.MkdirAll(complete, 0755)
	CopyDir(dst, src)

	completeDst := strings.Replace(complete, src, dst, -1)

	info, _ := os.Stat(completeDst)
	if !info.IsDirectory() {
		t.Errorf("CopyDir didnt' work properly, it didn't copy the source directory contents")
	}

	os.RemoveAll(complete)
	os.RemoveAll(completeDst)
}

