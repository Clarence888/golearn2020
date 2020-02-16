package main
//import "fmt"
import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/user/hello/morestrings"
)

func main(){
	fmt.Println(morestrings.ReverseRunes("Hello World"))
	fmt.Println(cmp.Diff("HelloWorld","HelloGO"))

}
