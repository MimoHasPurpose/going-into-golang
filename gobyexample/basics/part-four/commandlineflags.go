// flags are ways to specify options for command-line programs.
//  wc -l : -l is a flag.
// E:\Github\going-into-golang\gobyexample\basics\part-four>.\commandlineflags -word=opt -numb=7 -fork -svar=flag  
// E:\Github\going-into-golang\gobyexample\basics\part-four>commandlineflags -word=opt
// E:\Github\going-into-golang\gobyexample\basics\part-four>.\commandlineflags -word=opt a1 a2 a3
// E:\Github\going-into-golang\gobyexample\basics\part-four>.\commandlineflags -word=opt a1 a2 a3 -numb=7
// E:\Github\going-into-golang\gobyexample\basics\part-four>.\commandlineflags -wat                      


package main
import (
	"flag"
	"fmt"
)

func main(){
	wordPtr:=flag.String("word","foo","a string")
	numbPtr:=flag.Int("numb",42,"an int")
	forkPtr:=flag.Bool("fork",false,"a bool")
	var svar string
	flag.StringVar(&svar,"svar","bar","a string var")
	flag.Parse()
	fmt.Println("word: ",*wordPtr)
	fmt.Println("numb: ",*numbPtr)
	fmt.Println("fork: ",*forkPtr)
	fmt.Println("svar: ",svar)
	fmt.Println("tail: ",flag.Args())
}