package main

import (
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for{// 防止协程关闭
		buf := make([]byte,100) //创建缓存
		n,err :=conn.Read(buf)
		fmt.Println("来自客户端信息：",string(buf[:n]))
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write(buf[:n])
	}
}

func main(){
	// 创建socket数据结构
	listen,err := net.Listen("tcp","127.0.0.1:3000")
	if err != nil {
		fmt.Println("listen failed! msg:",err)
		return
	}
	fmt.Println("server 启动")

	for{ // 防止主进程启动后一次应答后退出
		conn,err := listen.Accept() //创建监听进程
		if err != nil {
			fmt.Println("accept failed")
		}
		// 将当前请求交到协程处理
		go handle(conn)
	}

}