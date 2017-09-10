package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

// md5All read all files content and calculate md5
func md5All(root string) (map[string][md5.Size]byte, error) {
	m := make(map[string][md5.Size]byte)
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("walk though path got a err: %v", err)
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		m[path] = md5.Sum(data)
		return nil
	}

	err := filepath.Walk(root, walkFn)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func main() {
	m, err := md5All(os.Args[1])
	if err != nil {
		fmt.Printf("md5All error: %v", err)
		return
	}

	var paths []string
	for path := range m {
		paths = append(paths, path)
	}

	sort.Strings(paths)

	for _, path := range paths {
		fmt.Printf("%x\t%s\n", m[path], path)
	}
}
