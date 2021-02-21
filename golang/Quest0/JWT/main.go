package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"encoding/base64"
	"encoding/json"
	"crypto/sha256"
	"crypto/hmac"
)
type structure struct{
	Typ string
	Alg string

	Sub string
	Exp string
	Iat string
	
}
type pltemp struct{
	Sub string
	Exp int
	Iat int
}
func jsonDecoding(data string)(structure,error){
	var jsonstruct structure
	err := json.Unmarshal([]byte(data),&jsonstruct)
	return jsonstruct,err
}
func hmac_sha256(data string,keystr []byte)(string){
	mac := hmac.New(sha256.New,keystr)
	mac.Write([]byte(data))
	res := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return res
}
func main(){
	file,err := os.Open("secret.txt")
	if err != nil{
		log.Fatal(err)
		return
	}
	key,err := ioutil.ReadAll(file)
	if err != nil{
		log.Fatal(err)
		return
	}
	header :="{\"typ\":\"JWT\",\"alg\":\"HS256\"}";
	hd := []byte(header)
	hdencoded := base64.StdEncoding.EncodeToString(hd)
	var payload pltemp
	payload.Sub = "Megumin"
	payload.Exp = 1612422922
	payload.Iat = 1612423222
	pl, err2 := json.MarshalIndent(payload,"","\t")
	if err2 != nil{
		log.Fatal(err2)
		return
	}
	plencoded := base64.StdEncoding.EncodeToString(pl)
	encodedstring := hdencoded + "." + plencoded
	signature := hmac_sha256(encodedstring,key)
	JWT := encodedstring + "." + signature
	fmt.Println(JWT)
	return
}
