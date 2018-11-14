package main

import (
	"fmt"
	"time"
)

func main() {
	if false { //周期
		ticker := time.NewTicker(time.Second * 5)
		go func() {
			for _ = range ticker.C {
				println("test")
			}
		}()

		time.Sleep(time.Minute)
	}

	if false { //倒计时
		timer1 := time.NewTimer(time.Second * 5) // 创建一个定时器
		<-timer1.C                               //等待信号

		println("test")

	}

	if false {
		timer2 := time.NewTimer(time.Second)
		go func() { //等触发时的信号

			<-timer2.C

			fmt.Println("Timer 2 expired")

		}() //由于上面的等待信号是在新线程中，所以代码会继续往下执行，停掉计时器

		time.Sleep(time.Second * 5)

	}

}
