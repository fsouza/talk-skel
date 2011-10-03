package main

import (
	"os"
	"testing"
)

func TestCopyDirectoryShouldCreateTheDestinationDirectory(t *testing.T) {
	src := "/tmp/src"
	dst := "/tmp/dst"
	os.MkdirAll(src, 0755)
	copyDir(dst, src)

	info, _ := os.Stat(dst)
	if !info.IsDirectory() {
		t.Errorf("CopyDir didn't created the destination directory, and it should!")
	}
}

