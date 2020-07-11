package osutil

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
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

func Rename(oldFile, newFile string) error {
	err := os.Rename(oldFile, newFile)
	if err != nil {
		err = CopyFile(oldFile, newFile)
		os.Remove(oldFile)
	}
	return err
}

func CreateTempFile(prefix string) (string, error) {
	file, err := ioutil.TempFile("", prefix)
	if err != nil {
		return "", CreateFileError
	}
	defer file.Close()
	return file.Name(), nil
}

func GetFileMd5(path string) (digest string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Open a new hash interface to write to
	hash := md5.New()
	// Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	// Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]
	// Convert the bytes to a string
	digest = hex.EncodeToString(hashInBytes)
	return digest, nil
}

func IsFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
