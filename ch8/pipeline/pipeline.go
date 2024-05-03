package main

import "fmt"

func counter(out chan<- int){
	for x:=0;x<6;x++{
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int){
	for x:= range in {
		out <- x*x
	}
	close(out)
}

func main(){
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)

	go squarer(squares, naturals)

	for x:= range squares{
		fmt.Println(x)
	}
}