package main

import (
	"fmt"
	"net"
)

func main()  {
	// 指定IP、port创建套接字
	conn,err := net.Dial("udp","127.0.0.1:8006")
	if err != nil {
		fmt.Println("net Dail err:",err)
		return
	}
	defer conn.Close() //main函数结束释放连接

	for i:=0;i<1000000;i++ {
		// 写数据到服务器
		conn.Write([]byte("Are you OK?"))
		//接受来自server数据
		buf :=make([]byte,4096)
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn read failed:",err)
			return
		}
		fmt.Println("客户端收到服务器数据：",string(buf[:n]))
	}
}
