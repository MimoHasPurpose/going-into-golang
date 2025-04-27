// flags are ways to specify options for command-line programs.
//  wc -l : -l is a flag.
package main
import (
	"fmt"
	"os"
)

func main(){
	wordPtr:=flag.String("word","foo","a string")
	numbPtr:=flag.Int("numb",42,"an int")
}