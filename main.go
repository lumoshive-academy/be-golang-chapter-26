package main

import (
	"fmt"
	"log"
)

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

	// Inisialisasi dengan InMemoryStorage
	inMemoryStorage, err := InitializeCachingData()
	if err != nil {
		log.Fatal(err)
	}
	inMemoryStorage.Save("Hello, In-Memory!")
	data, _ := inMemoryStorage.Load()
	fmt.Println(data) // Output: "Hello, In-Memory!"

	// Inisialisasi dengan DatabaseStorage
	databaseStorage, err := InitializeDatabaseStorage()
	if err != nil {
		log.Fatal(err)
	}
	databaseStorage.Save("Hello, Database!")
	data, _ = databaseStorage.Load()
	fmt.Println(data) // Output: "Hello, Database!"

	// InitializeApp dengan Notifier
	notifier, err := InitializeNotifier()
	if err != nil {
		log.Fatal(err)
	}

	notifier.SendNotification("Hello, this is a notification!")

	// binding value
	appConfig, err := InitializeAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	appConfig.PrintConfig()

	// struct field provider
	db := InitializeDatabase()
	fmt.Printf("Connected to database %s\n", db.Name)
}
