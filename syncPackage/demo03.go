package main

import (
	"fmt"
)

func test(a chan int,b chan int) {
	b <- 6
	b <- 7
	a <- 5
}

func main(){
	if false { // 非阻塞 channel先进先出,关闭channel 输入默认值 0
		var chanInt chan int = make(chan int,10)
		go func() {
			fmt.Println("开始gogo")
			defer close(chanInt)
			chanInt <- 1
			chanInt <- 2
			chanInt <- 3
			chanInt <- 4
		}()

		fmt.Println(<-chanInt)
		fmt.Println(<-chanInt)
		fmt.Println(<-chanInt)
		fmt.Println(<-chanInt)
		fmt.Println(<-chanInt)
		fmt.Println(<-chanInt)

	}

	if true {
		chan1 := make(chan int)
		chan2 := make(chan int,2)
		chanquit := 3
		go test(chan1, chan2)
		for {
			select {
			case n:=<-chan1:
				fmt.Println("chan1收到值",n)
				chanquit--
			case x := <-chan2:
				fmt.Println("chan2收到值", x)
				chanquit--
			default:
				if chanquit == 0 {
					return
				}
			}
		}

	}






}
