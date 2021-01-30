package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"structure"
	"encoding/json"
)
func httpGet(url string)(string, error){//返回一个包含着json的data字符串和一个get时候出现的error
	res, err := http.Get(url)//进行get操作
	if err != nil{
		fmt.Println("Error occurs when trying to 'get'")
	}//处理get得到的err
	defer res.Body.Close()//如果err了需要提前准备return时候关上文件
	data, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil{
		fmt.Println("Error occurs in ioutil");
	}//处理读res得到的err
	return string(data),err
}
func jsonDecoding(data string)(structure.bilibili,error){
	var jsonstruct structure.bilibili
	err := json.Unmarshal([]byte(data),&jsonstruct)
	return jsonstruct,err
}
func main(){
	jsonString, err1 := httpGet("https://api.live.bilibili.com/xlive/web-room/v1/index/getInfoByRoom?room_id=14866481")
	if err1 != nil{
		return
	}
	jsonStruct,err2 := jsonDecoding(jsonString)
	if err2 != nil{
		fmt.Println("Error occurs trying to decode json")
	}
	return
}