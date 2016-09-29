package main
import "fmt"
func receiverType(recv interface{}) {
	switch t:=recv.(type) {
	case *int:
		println("receiver is int")
		*(recv.(*int)) = 100500
	case *string:
		println("receiver is string")
		*(recv.(*string)) = "Yet another string"
	case *bool:
		println("receiver is bool")
		*(recv.(*bool)) = false
	case *float64:
		println("receiver is float")
		*(recv.(*float64)) = 2.718281828
	default:
		fmt.Printf("receiver defined as %#v\n", t)
	}
}

func main() {
	a := "Hello, playground"
	b := 1
	c := true
	d := 3.1415
	ch := make(chan struct{})
	e := &d
	receiverType(&a)
	receiverType(&b)
	receiverType(&c)
	receiverType(&d)
	receiverType(&ch)
	receiverType(&e)
	println(a)
	println(b)
	println(c)
	println(d)
}
