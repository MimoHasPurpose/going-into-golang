// 
package main
import (
	"flag"
	"fmt"
	"os"
)

func main(){
	fooCmd:=flag.NewFlagSet("foo",flag.ExitOnError)
	fooEnable:=fooCmd.Bool("enable",false,"enable")
	fooName:=fooCmd.String("name","","name")
	barCmd:=flag.NewFlagSet("bar",flag.ExitOnError)
}
