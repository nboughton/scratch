package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"math"
	"os"
	"path/filepath"

	"gopkg.in/cheggaaa/pb.v1"
)

// File stores a copy of the md5 hash and paths to all files that
// match that hash
type File struct {
	Hash  string
	Paths []string
}

// NewFile returns a new File with its hash and first found path
func NewFile(hash, path string) *File {
	return &File{
		Hash:  hash,
		Paths: []string{path},
	}
}

// AddPath adds a matching path to an existing File record
func (f *File) AddPath(path string) {
	f.Paths = append(f.Paths, path)
}

// Map files to hash of File objects
var h = make(map[string]*File)

func main() {
	// Check target dir
	dir := os.Args[1]
	fmt.Println("> Checking target directory.")
	if _, err := os.Stat(dir); err == os.ErrNotExist {
		fmt.Printf("No such directory [%s]\n", dir)
		os.Exit(1)
	}

	// Get count of all files to be checked
	fCount := 0
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			fCount++
		}

		return nil
	})
	fmt.Printf("> Files to process: %d\n", fCount)

	// Walk dir tree and populate hash
	fmt.Println("> Processing directory. This may take some time.")
	bar := pb.StartNew(fCount)
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			sum, err := md5Sum(path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Create or append record
			if _, ok := h[sum]; !ok {
				h[sum] = NewFile(sum, path)
			} else {
				h[sum].AddPath(path)
			}

			bar.Increment()
		}

		return nil
	})

	bar.Finish()

	// Iterate data to show duplicates
	fmt.Println("Report:")
	for _, r := range h {
		if len(r.Paths) > 1 {
			fmt.Printf("Duplicates found for %s\n", r.Paths[0])
			for _, p := range r.Paths[1:] {
				fmt.Println("\t", p)
			}
		}
	}
}

func md5Sum(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	s := bufio.NewScanner(f)
	for s.Scan() {
		if _, err := h.Write(s.Bytes()); err != nil {
			return "", err
		}
	}

	return string(h.Sum(nil)), nil
}
