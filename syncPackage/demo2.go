package main

import "fmt"

func main(){
	//channel

	ch := make(chan interface{})
	count := 2

	go func() {
		fmt.Println("Goroutine 1")
		ch <- "goroutine01 ending"
	}()

	go func() {
		fmt.Println("Goroutine 2")
		ch <- "goroutine02 ending"
	}()

	for range ch {
		count--
		if count == 0 {
			close(ch)
		}
		println(count)

	}
}
