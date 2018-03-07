package main

import "fmt"

type Something interface {
	DoThings(string, uint64) (int, error)
	DoNothing()
}

func main() {
	fmt.Println("vim-go")
}

func DoThings(string, uint64) (int, error) {
	return 1, nil
}

func DoNothing() {
	fmt.Println("vim-go")
}
