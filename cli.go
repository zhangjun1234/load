package main

import (
	"fmt"
	"os"
)

//这是一个用来接收命令行参数并且控制区块链操作的文件

type CLI struct {
}

const Usage = `命令提示 ：
	 .      "上传当前目录下所有文件,例如: upload . "
`

//接受参数的动作，我们放到一个函数中

func (cli *CLI) Run() {

	//1. 得到所有的命令
	args := os.Args
	if len(args) != 2 {
		fmt.Printf(Usage)
		return
	}

	//2. 分析命令
	cmd := args[1]
	switch cmd {
	case ".":
		fmt.Printf("上传文件\n")
		path,err:=os.Getwd()
		if err!=nil{
			fmt.Println(err)
			return
		}
		//tmpPath := strings.Split(AbsPath, "/")
		//path := tmpPath[len(tmpPath)-1]
		cli.UploadAllFile(path)
	default:
		fmt.Printf("无效的命令，请检查!\n")
		fmt.Printf(Usage)
	}
}

func (cli *CLI) UploadAllFile(path string) {
	UploadFiles(path)
}
