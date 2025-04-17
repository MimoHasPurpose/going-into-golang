
package main
import(
	"fmt"
	s "strings"
)
var P=fmt.Println

func main(){
	P("Contains: ", s.Contains("test","es"))
	P("To lower: ",s.ToLower("TEST"))
	P("ToUpper: ",s.ToUpper("test"))
	P("Index: ",s.Index("test","e"))
}