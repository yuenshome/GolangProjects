package main
import "fmt" //引入fmt库
import "regexp"
import "reflect"
func main() {
	fmt.Println("Method 1")
	s := "Hello/?keywords=xiaofei"
	reg := regexp.MustCompile(`\?.*`)
	//fmt.Println(reg.FindString("Hello/?keywords=xiaofei"))
	fmt.Println(reg.FindString(s))
	fmt.Println("Method 2")
	reg = regexp.MustCompile(`keywords.*`)
	fmt.Println(reg.FindString("Hello/?keywords=xiaofei"))
	fmt.Println("type:", reflect.TypeOf(s))
}
