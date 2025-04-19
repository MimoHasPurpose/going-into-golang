package main
import (
	b64 "encoding/base64"
	"fmt"
)
func main(){


	data:="abc"
	sEnc:=b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	data2:="hello world"
	s:=b64.StdEncoding.EncodeToString([]byte (data2))
	fmt.Println(s)
	sDec,_:=b64.StdEncoding.DecodeString(s)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc:=b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec,_:=b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}