// channels are pipe that connect concurrent goroutines.
// u can send values into channels from on goroutine and receive those values into another goroutine.

package main
import "fmt"
func main(){
	messages:=make(chan string)		// create a new channel by make.

	tommy:=make(chan string)
	go func(){tommy<-"hello tommy"}()
	go func(){ messages<- "ping"}()

	msg:=<-messages
	tom:=<-tommy
	fmt.Println(tom)
	fmt.Println(msg)
}