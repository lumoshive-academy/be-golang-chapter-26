package main

import "fmt"

func main() {
	service := InitializMyService()
	fmt.Println(service.Greeter.Message)
}
