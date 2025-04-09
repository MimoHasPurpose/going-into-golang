// errors are the last return value and have type error, a built-in interface 
// 

package main
import (
	"errors"
	"fmt"
)

func f(arg int)(int , error){
	if arg==42{
		return -1,errors.New("cant work with 42")
	}
	return arg+3,nil
}

var ErrOutOfTea=fmt.Error("no more tea availabe")
var ErrPower=fmt.Errorf("can't boil water")

func makeTea(arg int) error{
	if arg==2{
		return ErrOutOfTea
	}else if arg==4{
		return fmt.Errorf("making tea: %w",ErrPower)
	}
	return nil
}


func main(){
	for _,i:=range[]int{7,42}{
		if r,e:=f(i); e!=nil{
			fmt.Println("f failed: ",e)
		}else{
			fmt.Println("f worked: ",r)
		}
	}

	for i:=range 5{
		if err:=makeTea(i); err!=nil{
			if error.Is(err,ErrOutOfTea){
				fmt.Println("We should bu new tea! ")
			}else if errors.Is(err,ErrPower){
				fmt.Println("Now it is dark. ")
			}else {
				fmt.Printf("unknown error: %s \n",err)
			}
			continue
		}
		fmt.Println("Tea is ready! ")
	}


}





