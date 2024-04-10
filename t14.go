package main

import(
	"fmt"
	"reflect"
)

// 1. Using reflect
func ref(){
	ch := make(chan int)
	arr := []interface{}{42, "hello", false, ch}

	for _, v := range arr {
		v := reflect.ValueOf(v)
		fmt.Printf("%v is of type %s \n", v, v.Kind().String())
	}
}

// 2. Using switch
func swi(){
	ch := make(chan int)
	arr := []interface{}{42, "hello", false, ch, struct{}{}}

	for _, v := range arr {
		switch valueType := v.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("bool")
		default:
			fmt.Printf("%T\n", valueType)
		}
	}
}

func main() {
	ref()
	swi()
}
