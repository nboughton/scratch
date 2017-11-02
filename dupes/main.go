package main

import (
	"crypto/md5"
	"fmt"
	"io"
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

func main() {
	dir := os.Args[1]
	if _, err := os.Stat(dir); err == os.ErrNotExist {
		fmt.Printf("No such directory [%s]\n", dir)
		os.Exit(1)
	}

	fmt.Println("Preparing. This may take some time.")
	data, errors := genData(dir)
	for _, r := range data {
		if len(r.Paths) > 1 {
			fmt.Printf("\nDuplicates found for %s\n", r.Paths[0])
			for _, p := range r.Paths[1:] {
				fmt.Println("\t", p)
			}
		}
	}

	fmt.Println("The following errors occurred during the run:")
	for _, err := range errors {
		fmt.Println(err)
	}
}

func genData(dir string) (map[string]*File, []error) {
	h, errors, bar := make(map[string]*File), []error{}, pb.StartNew(countFiles(dir))

	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			sum, err := hashSum(path)
			if err != nil {
				errors = append(errors, err)
				bar.Increment()
				return nil
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

	return h, errors
}

func countFiles(dir string) int {
	c := 0

	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			c++
		}

		return nil
	})

	return c
}

func hashSum(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return string(h.Sum(nil)), nil
}
