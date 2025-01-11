package main

import(
	"fmt"
	"time"
)
func main(){
	i:=2
	fmt.Print("Write ", i," as ")
	switch(i){
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}
	a:=time.Now()
	fmt.Println("time taken to execute this.")
	fmt.Println("time taken to execute this.")
	fmt.Println("time taken to execute this.")
	fmt.Println("time taken to execute this.")
	fmt.Println("time taken to execute this.")
	fmt.Println("time taken to execute this.")
	b:=time.Now()
	c:=b.Sub(a)
	fmt.Println(c)


}