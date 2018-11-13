package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var locker = new(sync.Mutex) // 创建一把锁
var cond = sync.NewCond(locker) // 条件变量是构建在一个基础锁上的同步原语
var capacity = 10
var consumerNum = 3 //消费者
var producerNum = 5 //生产者

func producer(out chan<- int) {
	for i:= 0;i < producerNum ; i++ {
		go func(nu int) {
			for {
				cond.L.Lock() //对竞争资源（边界值上锁)
				for len(out)  == capacity { // 如果满了，循环等待
					fmt.Println("Capacity Full, stop Produce")
					cond.Wait() // 当前goroutine 为阻塞态
				}
				num := rand.Intn(100)
				out <- num //生产
				fmt.Printf("Gouroutine %d producer:num %d \n",nu,num)
				cond.L.Unlock() //释放锁
				cond.Signal() // 通知阻塞态gouruntine更改为就绪态
				time.Sleep(time.Second)
			}
		}(i)
	}
}

func consumer(in <-chan int) {
	for i:=0;i< consumerNum;i++ {
		go func (nu int) {
			for {
				fmt.Println("producer get lock")
				cond.L.Lock() //临界区加锁
				fmt.Println("producer get lock0")

				for len(in) == 0 {
					fmt.Println("capacity Empty,stop Cunsume")
					fmt.Println("producer get lock1")
					cond.Wait() // 阻塞当前 gouroutine，次数锁还是在当前gouroutine手中，等待其他gouroutine释放signal
					fmt.Println("producer get lock2")
				}
				fmt.Println("producer get lock3")

				num := <-in
				fmt.Printf("Gouroutine %d: consume num %d \n",nu,num)
				cond.L.Unlock()
				time.Sleep(time.Microsecond * 500 )
				cond.Signal() //通知其他阻塞 gouroutine
			}
		}(i)
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	quit := make(chan bool)
	product := make(chan int, capacity)

	producer(product)
	consumer(product)

	<-quit
}
