package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/nboughton/go-utils/input"
	"github.com/nboughton/go-utils/regex/common"
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

// Keep removes the files at all paths except for the 1 specifed by its index number
func (f *File) Keep(idx int) error {
	if idx > len(f.Paths) {
		return fmt.Errorf("Error: Invalid index: [%d]", idx)
	}

	for i := range f.Paths {
		if i != idx {
			fmt.Println("Removing: ", f.Paths[i])
			if err := os.Remove(f.Paths[i]); err != nil {
				return err
			}
		}
	}

	return nil
}

// Index returns the index values of all paths for file
func (f *File) Index() (idx []string) {
	for i, p := range f.Paths {
		idx = append(idx, fmt.Sprintf("[%d] %s", i, p))
	}

	return idx
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
			fmt.Printf("\nDupes found for %s\n", r.Paths[0])
			for _, p := range r.Index() {
				fmt.Println("\t", p)
			}

			fmt.Print("Remove dupes? [Y/n or index of file to keep]: ")
			ans := input.ReadLine()
			idx, err := strconv.Atoi(ans)
			if err != nil && !common.Yes.MatchString(ans) && ans != "" {
				fmt.Println("No action taken. Continuing.")
				continue
			}

			if err := r.Keep(idx); err != nil {
				errors = append(errors)
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	if len(errors) > 0 {
		fmt.Println("\nThe following errors occurred during the run:")
		for _, err := range errors {
			fmt.Println(err)
		}
	}
}

func genData(dir string) (map[string]*File, []error) {
	h, errors, bar := make(map[string]*File), []error{}, pb.StartNew(countFiles(dir))

	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if validFile(f) {
			sum, err := hashSum(path)
			if err != nil {
				errors = append(errors, err)
				fmt.Println(err)
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
			time.Sleep(time.Millisecond)
		}

		return nil
	})
	bar.Finish()

	return h, errors
}

func countFiles(dir string) int {
	c := 0

	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if validFile(f) {
			c++
		}

		return nil
	})

	return c
}

func validFile(f os.FileInfo) bool {
	if !f.IsDir() && f.Size() > 0 && !strings.HasPrefix(f.Mode().String(), "L") {
		return true
	}

	return false
}

func hashSum(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return string(h.Sum(nil)), nil
}
