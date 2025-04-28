package  main
import(
	"fmt"
	// "bufio"
	// "io"
	"os"
)

func check(e error){
	if e!=nil{
		panic(e)
	}
}


func  main(){
	dat,err:=os.ReadFile("E:\\Github\\going-into-golang\\projects\\temp\\data.txt")
	check(err)
	fmt.Println(string(dat))
	code,err:=os.ReadFile("E:\\Github\\going-into-golang\\projects\\filemanipulation.go")
	check(err)
	fmt.Println(string(code))
	f,err:=os.Create("E:\\Github\\going-into-golang\\projects\\temp\\data.txt")
	check(err)
	defer f.Close()
	n2,err:=f.Write(code)
	check(err)
	fmt.Printf("wrote %d bytes \n",n2)
}