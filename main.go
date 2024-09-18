package main

import "fmt"

func main() {
	// service := InitializMyService()
	// fmt.Println(service.Greeter.Message)

	// service, err := InitializMyService()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println(service.Greeter.Message)

	service, err := InitializMyService("Lumoshive")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(service.Greeter.Message)

	config, err := InitializeServiceConfig()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(config.ConfigA.Message)
}
