// we use channels to sync execution across goroutines. 
// when waiting for multiple goroutines to finish you may prefer to use a waitgroup.


package main
import (
	"fmt"
	"time"

)

func worker(done chan bool){
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done<-true	
}

func main(){
	done :=make(chan bool,1)
	go wor/ker(done)
	<-done
}