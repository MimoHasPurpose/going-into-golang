package main

import(
	"fmt"
	"bufio"
	"os"

)

func check(e error){
	if e!=nil{
		panic(e)
	}
}


func main(){

	d1:=[]byte("hello/ngo/n")
	err:=os.WriteFile("E:/Github/going-into-golang/gobyexample/basics/part-four/tmp/dat1.txt",d1,0644)
	check(err)
	f,err:=os.Create("E:/Github/going-into-golang/gobyexample/basics/part-four/tmp/dat1")

	check(err)
	defer f.Close()

	d2:=[]byte{115,111,109,101,10}
	n2,err:=f.Write(d2)

	check(err)
	fmt.Printf("wrote %d bytes \n",n2)

	n3,err:=f.WriteString("Hello bollo babuabababa\n")
	check(err)
	fmt.Printf("wrote %d bytes \n",n3)

	f.Sync()

	w:=bufio.NewWriter(f)
	n4,err:=w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes \n",n4)
	w.Flush()



}