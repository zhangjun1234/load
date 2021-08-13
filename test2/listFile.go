package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const urlstring = "http://127.0.0.1:5001/api/v0/add?recursive=true"
var filePaths [] string
func UploadFiles(path string)  {
	//listFiles("/home/miaowu/test","http://127.0.0.1:5001/api/v0/add")
	listFiles(path)
	fmt.Println("==============================================")
	for _,fp :=range filePaths{
		resp,err :=PostFile(fp,urlstring)
		if err!=nil{
			fmt.Println("err : ",err)
			return
		}
		body,err:=ioutil.ReadAll(resp.Body)
		if err!=nil {
			fmt.Println("read err :",err)
			return
		}
		fmt.Println(string(body))
	}
}

func listFiles(dirname string) {
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil{
		log.Panic(err)
	}
	for _, fi := range fileInfos {
		filename := dirname + "/" + fi.Name()
		if fi.IsDir() {
			//继续遍历fi这个目录
			listFiles(filename)
		}else{
			filePaths = append(filePaths, filename)
		}
	}
}
