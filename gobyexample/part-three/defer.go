// defer is used to ensure that a function call is performed later in a program's execution. 
// 


package main
import(
	"fmt"
	"os"
)
func main(){
	f:=createFile("E:/Github/going-into-golang/gobyexample/part-three/defer/defer.txt")
	defer closeFile(f)
	writeFile(f)

}

func createFile(p string) *os.File{
	fmt.Println("creating")
	f,err:=os.Create(p)
	if err!=nil{
		panic(err)
	}
	return f
}

func writeFile(f *os.File){
	fmt.Println("writing")
	fmt.Fprintln(f,"data")

}
func closeFile(f *os.File){
	fmt.Println("closing")
	err:=f.Close()
	if err!=nil{
		panic(err)
	}
}
