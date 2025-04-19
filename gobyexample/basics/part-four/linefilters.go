//  program that reads i/p on stdin process it then prints some derived results to stdout
// grep and sed are common line filters


package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"

)


func main(){
	scanner:=bufio.NewScanner(os.Stdin)


	for scanner.Scan(){
		ucl:=strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	if err:=scanner.Err(); err!=nil{
		fmt.Fprintln(os.Stderr,"error: ",err)
		os.Exit(1)
	}
}