package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
	res := make(chan result)
	errc := make(chan error, 1)

	go func() {
		var wg sync.WaitGroup

		walkFn := func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("walk though path got a err: %v", err)
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			wg.Add(1)
			go func() {
				data, err := ioutil.ReadFile(path)
				select {
				case res <- result{path, md5.Sum(data), err}:
				case <-done:
				}

				wg.Done()
			}()

			select {
			case <-done:
				return fmt.Errorf("walk canceled")
			default:
				return nil
			}
		}

		err := filepath.Walk(root, walkFn)
		go func() {
			wg.Wait()
			close(res)
		}()

		errc <- err
	}()

	return res, errc

}

// md5All read all files content and calculate md5
func md5All(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	results, errc := sumFiles(done, root)

	m := make(map[string][md5.Size]byte)
	for r := range results {
		if r.err != nil {
			fmt.Printf("get error in result from results channel: %v", r.err)
			return nil, r.err
		}
		m[r.path] = r.sum
	}
	if err := <-errc; err != nil {
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
