package main
import(
	"fmt"
	"time"
)

func main(){
	p:=fmt.Println
	now:=time.Now()
	p(now)
	then:=time.Date(
		2009,11,17,20,34,58,651387237,time.UTC)
	p(then)

	p(then.Year(),then.Month(),then.Day())
	diff:=now.Sub(then)
	p(diff)

	p(diff.Hours())
}