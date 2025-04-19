package main
import(
	"bufio"
	"fmt"
	"io"
	"os"

)


func check(e error){
	if e!=nil{
		panic(e)
	}
}


func main(){
	dat,err:=os.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(dat))

	f,err:=os.Open("/tmp/dat")
	check(err)


	b1:=make([]byte,5)
	n1,err:=f.Read(b1)

	check(err)
	fmt.Printf("%d bytes: %s \n",n1,string(b1[:n1]))


	b1:=make([] byte,5)
	n1,err:=f.Read(b1)
	




}