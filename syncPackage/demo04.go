package main

import "sync"

func main(){
	var wg sync.WaitGroup // 一组走一波，不抛弃不放弃
	wg.Add(2) //上俩个兄弟，不抛弃不放弃，标记 为2

	go func() {
		println("goroutine 1")
		wg.Done() // 完成
	}()
	go func() {
		println("gouroutine 2")
		wg.Done() //gogo
	}()
	println("都到齐了吗")
	wg.Wait() //等待兄弟归来
	println("兄弟归来")
}
