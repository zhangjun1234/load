package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const urlstring = "http://127.0.0.1:5001/api/v0/add?recursive=true"
	resp,err :=PostFolder("hello",urlstring)
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
