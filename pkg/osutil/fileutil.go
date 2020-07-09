package main

import (
	"io"
	"os"

	"github.com/pkg/errors"
)

var (
	CreateFileError = errors.New("failed to create file")
	OpenFileError   = errors.New("failed to open file")
	InvalidTarget   = errors.New("invalid target specified")
)

func TouchFile(name string) (err error) {
	empty, err := os.Create(name)
	if err != nil {
		return
	}
	defer empty.Close()
	return nil
}

func ReadFile() error {
	return nil
}

func WriteFile() error {
	return nil
}

func CopyFile(src, dst string) (err error) {
	// check if source file is a regular file
	instat, err := os.Stat(src)
	if err != nil {
		return
	} else if !instat.Mode().IsRegular() {
		return InvalidTarget
	}

	// copy infile contents to outfile
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Open(dst)
	if err != nil {
		return
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	if err = os.Chtimes(dst, instat.ModTime(), instat.ModTime()); err != nil {
		return
	}

	return nil
}

func Rename() error {
	return nil
}

func GetFileMd5() string {
	return ""
}

// TODO: Touch/Read/Write/Copy/Rename Temp File
