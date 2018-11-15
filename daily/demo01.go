package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main(){
	if false { //flag命令行参数解析
		married := flag.Bool("married",false,"Are you married ?")
		age := flag.Int("age",22,"How old are you ?")
		name := flag.String("name","everyone","what is your name")
		var address string
		flag.StringVar(&address,"address","where is your address ?","the greeting object")
		flag.Parse() // 解析输入参数

		fmt.Println("输出的参数married的值是:", *married)//不加*号的话,输出的是内存地址
		fmt.Println("输出的参数age的值是:", *age)
		fmt.Println("输出的参数name的值是:", *name)
		fmt.Println("输出的参数address的值是:", address)
	}

	if false {
		cmd := exec.Command("go","version")
		fmt.Println(cmd.Args,cmd.Path)
	}

	if false {
		cmd := exec.Command("cat")
		stin,err := cmd.StdinPipe() // stin连接stdPipe

		if err != nil {
			fmt.Println(err)
		}

		_,err = stin.Write([]byte("tmp.txt"))
		if err != nil {
			println(err)
		}

		stin.Close()
		cmd.Stdout = os.Stdout
		cmd.Start()
	}

	if true {
		cmd := exec.Command("ls")
		stdout,_ :=cmd.StdoutPipe()
		cmd.Start()
		content,_ := ioutil.ReadAll(stdout)
		fmt.Println(string(content))

	}






}
