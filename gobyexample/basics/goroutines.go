// lightweight thread of execution
// to invoke a coroutine: go f(s)
package main
import(
	"fmt"
	"time"
)

func f(from string){
	for i:=0;i<30;i++{
		fmt.Println(from,":",i)
	}
}
func ff(from string){
	for i:=0;i<1;i++{
		fmt.Println(from,": ",i)
	}
}

func main(){
	f("direct")
	
	time.Sleep(time.Second)
	go ff("goroutine")

	go func(msg string){
		fmt.Println(msg)

	}("going")
	// time.Sleep(time.Second)
	// time.Sleep(time.Second)
	fmt.Println("done")
}