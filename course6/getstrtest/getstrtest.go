package getstrtest
import "fmt"
var str string = "hello world in getstrtest."

func init() {
	fmt.Println("Initialize getstrtest package.")
}
func getstring() string {
	return str
}

func Getstring() string {
	return str
}