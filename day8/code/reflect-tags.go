package main

import "reflect"

type A struct {
	B string `k:"this" orm:"column_b"`
	C int	 `k:"that"`
}

func main() {

	a := &A{}

	//field, _ := reflect.TypeOf(a).Elem().FieldByName("B")

	for i := 0; i < 10; i++ {

		f := reflect.TypeOf(a).Elem().Field(i)
		if f.Name == "" {
			break
		}
		println(string(f.Name), string(f.Tag))

	}
}