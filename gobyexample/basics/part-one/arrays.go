package main
import "fmt"

func main(){
	var a[5]int
	fmt.Println("emp:", a)

	a[4]=100
	fmt.Println("set: ",a)
	fmt.Println("get: ",a[4])

	fmt.Println("len: ",len(a))

	// b:=[5]int{1,2,3,4,5}
	// fmt.Println("dcl: ",b)

	b:=[...]int{2,3,4,1,5,2,3}
	fmt.Println("dcl: ",b)
	c:=[...]int{100,10:400,500}
	// if u specify index with : , element in between will be zeroed
	fmt.Println("idx: ",c)

	var twoD [2][3]int
	for  i:=0;i<2;i++{
		for j :=0;j<3;j++{
			twoD[i][j]=i+j
		}
	}
	fmt.Println("2d: ", twoD)

	threeD:=[3][3]int{	
		{1,1,1},
		{2,2,2},
		{3,3,3},
	}
	fmt.Println("3x3 array: ",threeD)
}