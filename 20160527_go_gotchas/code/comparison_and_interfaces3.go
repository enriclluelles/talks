package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

//START COMP OMIT
func main() {
	err := zipDecompress("/tmp/gopher.zip", "gopher.png")
	if err != nil {
		fmt.Println("Back in main, error is not nil")
		switch derr := err.(type) {
		case *DecompressingError:
			log.Fatalf("failure while decompressing %s with method %d: %s\n", derr.Name, derr.Method, derr)
		default:
			log.Fatalf("darn, it failed %v\n", err)
		}
	}
}

type DecompressingError struct {
	error
	zip.FileHeader
}

//END COMP OMIT

func findFileInZip(zipFile, file string) (*zip.File, error) {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return nil, err
	}
	for _, f := range reader.File {
		if f.Name == file {
			return f, nil
		}
	}
	return nil, fmt.Errorf("Couldn't find' file %s", file)
}

//START COMP2 OMIT
func zipDecompress(zipFile, file string) error {
	zp, err := findFileInZip(zipFile, file)
	if err != nil {
		return err
	}

	nf, err := os.Create(path.Join(path.Dir(zipFile), zp.Name))
	if err != nil {
		return &DecompressingError{err, zp.FileHeader}
	}

	fo, err := zp.Open()
	if err != nil {
		return &DecompressingError{err, zp.FileHeader}
	}

	_, err = io.Copy(nf, fo)

	if err != nil {
		return &DecompressingError{err, zp.FileHeader}
	}
	return nil
}

//END COMP2 OMIT
