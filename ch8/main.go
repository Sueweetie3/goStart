package main

import (
	"ch8/thumbnail"
	"os"
	"sync"
)

func makeThumbnails(filenames []string){
	ch := make(chan struct{})
	for _, f := range filenames{
		go func(f string){
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}
	for range filenames{
		<-ch
	}
}

func makeThumbnails2(filenames []string)(thumbfiles []string, err error){
	type item struct{
		thumbfile string
		err error
	}
	ch := make(chan item, len(filenames))
	for _, f := range filenames{
		go func(f string){
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames{
		it := <-ch
		if it.err != nil{
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

func makeThumbnails3(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames{
		wg.Add(1)
		go func(f string){
			defer wg.Done()
			thumbfile, err := thumbnail.ImageFile(f)
			if err != nil{
				return;
			}
			info, _ := os.Stat(thumbfile)
			sizes <- info.Size()
		}(f)
	}

	go func(){
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes{
		total += size
	}
	return total
}

func main(){
	// makeThumbnails([]string{"./pic/001.jpg", "./pic/002.jpg", "./pic/003.jpg"})
	// thumbfiles, err := makeThumbnails2([]string{"./pic/001.jpg", "./pic/002.jpg", "./pic/003.jpg"})
	// if err != nil{
	// 	panic(err)
	// }
	// for _, t := range thumbfiles{
	// 	println(t)
	// }
	var ch = make(chan string)
	go func(){
		ch <- "./pic/001.jpg"
		ch <- "./pic/002.jpg"
		ch <- "./pic/003.jpg"
		close(ch)
	}()
	sizes := makeThumbnails3(ch)
	println("Total size:", sizes)
}