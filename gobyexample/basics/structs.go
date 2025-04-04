// structs are typed collection of fields. 
// group data to form record
// go is garbage collected language: safely return a pointer to a local variable: only cleaned by garbage collector when no active references.
// 

package main
import "fmt"
type person struct{
	name string
	age int
}
func newPerson(name string) *person{
	p:=person{name:name}
	p.age=42
	return &p
}
func main(){
	fmt.Println(person{"Bob",20})
	fmt.Println(person{name:"Alice",age:30})
	fmt.Println(person{name:"Fred"})
	fmt.Println(person{name:"Fred"})	//ommited fields will be zero valued.
	fmt.Println(&person{name:"Ann",age:40}) // & prefix yields a pointer to the struct.
	fmt.Println(newPerson("jon"))


	s:=person{name:"sanjay",age:444}
	fmt.Println(s.name)

	sp:=&s
	fmt.Println(sp.age)

	dog :=struct{
		name string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

}
