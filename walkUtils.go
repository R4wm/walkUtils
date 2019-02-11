package walkUtils

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// PrintMD5Sum Just a filewalk exercise, maybe update and use later.
func PrintMD5Sum(path string, info os.FileInfo, err error) error {

	// Handle error
	if err != nil {
		log.Print(err)
		return nil
	}

	if info.IsDir() {
		dir := filepath.Base(path)
		log.Printf("Skip: %s is a directory -> %s\n", path, dir)
		return nil

	} else if info.Size() == 0 {
		log.Printf("Skip: Not a dir but size is 0, so skipping %s.", path)
		return nil

	} else if !info.Mode().IsRegular() {
		log.Printf("Skip: A Symlink: %s -> %s\n", path, info.Mode())
		return nil

	} else {
		f, err := os.Open(path)
		if err != nil {
			return err
		}

		defer f.Close()

		h := md5.New()
		// copy f contents into h
		if _, err := io.Copy(h, f); err != nil {
			log.Printf("Skip: From ls() failed to copy contents into h")
			log.Fatal(err)
		}

		fmt.Printf("%x %s\n", h.Sum(nil), path)
	}
	return nil
}
