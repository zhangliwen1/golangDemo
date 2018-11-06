package main

import (
	"fmt"
	"net"
	"time"
)

func main(){

	// 创建一个udp地址结构，指定服务器IP+port
	serverAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:8006")
	if err != nil {
		fmt.Println("ResolverUDPAddr err:",err)
		return
	}

	//创建socket通信
	udpConn,err := net.ListenUDP("udp",serverAddr)
	if err!= nil {
		fmt.Println("ListenUDP err:",err)
		return
	}

	defer udpConn.Close()
	//创建服务器通信socket完成

	// 创建变量接受客户端发送过来数据
	buf :=make([]byte,4096) // 1024 *4,4k

	for {
		n,cltAddr,err := udpConn.ReadFromUDP(buf) //读取字节数、客户端地址、error
		if err != nil {
			fmt.Println("Read From Udp err:",err)
			return
		}
		// 处理数据
		fmt.Printf("服务器读取到 %v 的数据：%s",n,string(buf[0:n]))

		// 启动goroutine处理响应
		go func (){
			// 获取当前系统时间
			daytime := time.Now().String()

			_,err := udpConn.WriteToUDP([]byte(daytime),cltAddr)
			if err != nil {
				fmt.Println("Write to UDP error:",err)
				return
			}

		}()
	}
}
