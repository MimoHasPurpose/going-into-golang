package main 
import(
	"fmt"
	// "net"
	"net/url"	
)
func main(){
	s:="postgres://user:pass@host.com:5432/path?k=v#f"
	u,err:=url.Parse(s)
	if err!=nil{
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User,u.User.Username())
	p,_:=u.User.Password()
	fmt.Println(p)
	// fmt.Println(host)

}