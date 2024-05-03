package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64){
	defer wg.Done()
	if cancelled(){
		return;
	}
	for _, entry := range dirents(dir){
		if entry.IsDir(){
			wg.Add(1)
			go walkDir(filepath.Join(dir, entry.Name()), wg, fileSizes)
		} else {
			fileSizes<- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 10)
func dirents(dir string) []os.FileInfo{
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func(){<-sema}()
	entries, err :=  ioutil.ReadDir(dir)
	if err != nil{
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

var done = make(chan struct{})
func cancelled() bool{
	select {
	case <-done:
		return true
	default: 
		return false
	}
}

func printDiskUsage(nfiles, nbytes int64) {
    fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main(){
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var tick <-chan time.Time
	if *verbose{
		tick = time.Tick(500 * time.Millisecond)
	}

	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	go func(){
		for _, root := range roots{
			wg.Add(1)
			go walkDir(root, &wg, fileSizes)
		}
	}()

	go func(){
		wg.Wait()
		close(fileSizes)
	}()

	go func(){
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	var nfiles, nbytes int64
	loop:
		for{
			select {
			case <-done:
				for range fileSizes {

				}
				return;
			case size, ok := <-fileSizes:
				if !ok{
					break loop
				}
				nfiles++
				nbytes += size
			case <-tick:
				printDiskUsage(nfiles, nbytes)
			}
		}
	printDiskUsage(nfiles, nbytes)
}