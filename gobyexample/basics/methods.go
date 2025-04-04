// go supports methods defined on struct types.

package main
import "fmt"

type rect struct {
	width,height int
}
 func (r *rect) area() int{	//pointer receiver
	return r.width*r.height
 }
  
 func (r rect) perim() int{	//value receiver
	return 2*r.width+2*r.height
 }

 func main(){

	r:=rect{width:10,height:5}
	fmt.Println("area: ",r.area())
	fmt.Println("area: ",r.perim())
	// fmt.Println("perim: ",r.perim())
	rp:=&r
	fmt.Println("area: ",rp.area())
	fmt.Println("perim: ",rp.perim())



 }