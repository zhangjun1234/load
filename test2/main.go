package main

import (
	"fmt"
	"io/ioutil"
)

const target_url = "http://127.0.0.1:5001/api/v0/add?chunker=size-262144&encoding=json&hash=sha2-256&inline-limit=32&pin=true&progress=true&recursive=true&stream-channels=true"
func main(){
	//path,err:=os.Getwd()
	//if err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(path)
	//a := strings.Split(path, "/")
	//fmt.Println(a[len(a)-1])
	resp,err := PostFile("/home/miaowu/test/1.txt",target_url)
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