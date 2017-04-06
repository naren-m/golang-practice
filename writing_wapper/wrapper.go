// http://chrisarges.net/2016/05/20/creating-a-function-wrapper-in-golang.html
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func callme(fn interface{}, params ...interface{}) (result []reflect.Value) {
	f := reflect.ValueOf(fn)
	if f.Type().NumIn() != len(params) {
		panic("incorrect number of parameters!")
	}
	inputs := make([]reflect.Value, len(params))
	for k, in := range params {
		inputs[k] = reflect.ValueOf(in)
	}
	return f.Call(inputs)
}

func log(fn interface{}, params ...interface{}) (result []reflect.Value) {
	f := reflect.ValueOf(fn)
	if f.Type().NumIn() != len(params) {
		panic("incorrect number of parameters!")
	}
	inputs := make([]reflect.Value, len(params))
	for k, in := range params {
		inputs[k] = reflect.ValueOf(in)
	}
	return f.Call(inputs)
}

func hello(i int) {
	fmt.Println("hello " + strconv.Itoa(i))
}

func hiya(name string) {
	fmt.Println("hiya " + name)
}

func awesome(i int, name string) {
	fmt.Println("high " + strconv.Itoa(i) + ", " + name)
}

func info(name string) {
	fmt.Println("info " + name)
}

func main() {
	callme(hello, 1)
	callme(hiya, "buddy")
	callme(awesome, 5, "dude")

	log(info, "Logging info")
}
