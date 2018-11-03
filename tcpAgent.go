package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main(){
	conn,err := net.Dial("tcp","127.0.0.1:3000")
	if err!= nil {
		fmt.Println("err dialing:",err.Error())
		return
	}
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin) // 创建一个标准输入器
	for {
		str,_ := inputReader.ReadString('\n')
		data := strings.Trim(str,"\n")
		if data == "quit" {
			return
		}
		_,err := conn.Write([]byte(data)) //将请求数据发送给服务器
		if err != nil {
			fmt.Println("send data error:",err)
			return
		}
		//接受来自客户数据
		buf := make([]byte,512)
		n,_ :=conn.Read(buf)
		fmt.Println("from server:",string(buf[:n]))
	}
}