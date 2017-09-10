/*bounded parallelism

Author: Google & zouying

Three stage:
1. walk the tree,
2. read and digest the files,
3. and collect the digests

*/
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

// walkFiles starts a goroutine to walk the directory tree at root and send the
// path of each regular file on the string channel.  It sends the result of the
// walk on the error channel.  If done is closed, walkFiles abandons its work.
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("walk though path got a err: %v", err)
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		select {
		case paths <- path:
		case <-done:
			return fmt.Errorf("walk canceled")
		}

		return nil
	}

	go func() {
		defer close(paths)
		errc <- filepath.Walk(root, walkFn)
	}()

	return paths, errc
}

// digester reads path names from paths and sends digests of the corresponding
// files on c until either paths or done is closed.
func digester(done <-chan struct{}, paths <-chan string, results chan<- result) {
	for path := range paths {
		data, err := ioutil.ReadFile(path)
		select {
		case results <- result{path, md5.Sum(data), err}:
		case <-done:
			return
		}
	}
}

// MD5All reads all the files in the file tree rooted at root and returns a map
// from file path to the MD5 sum of the file's contents.  If the directory walk
// fails or any read operation fails, MD5All returns an error.  In that case,
// MD5All does not wait for inflight read operations to complete.
func md5All(root string) (map[string][md5.Size]byte, error) {

	done := make(chan struct{})
	defer close(done)
	paths, errc := walkFiles(done, root)

	res := make(chan result)
	var wg sync.WaitGroup
	const numDigesters = 20
	wg.Add(numDigesters)
	for i := 0; i < numDigesters; i++ {
		go func() {
			digester(done, paths, res)
			wg.Done()
		}()
	}

	// wait all digesters finished, and close result channel
	go func() {
		wg.Wait()
		close(res)
	}()

	m := make(map[string][md5.Size]byte)
	for r := range res {
		if r.err != nil {
			fmt.Printf("digester the file(%s) is error: %v", r.path, r.err)
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
