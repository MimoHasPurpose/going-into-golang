// SHA256 hashes used to compute short identities for text blobs.
package main
import(
	"crypto/sha256"
	"fmt"
)


func main(){
	s:="sha256 this string"
	h:=sha256.New()
	// string coerced to bytes.
	h.Write([]byte(s))
	bs:=h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n",bs)
}