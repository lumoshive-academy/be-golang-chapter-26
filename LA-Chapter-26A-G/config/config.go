// package config

// type Config struct {
// 	Message string
// }

// func NewConfig() *Config {
// 	return &Config{Message: "Config 1"}
// }

// func NewConfigAlternative() *Config {
// 	return &Config{Message: "Config 2"}
// }

package config

import "fmt"

type Config struct {
	Message string
}

type ConfigA Config
type ConfigB Config

func NewConfig() *ConfigA {
	config := &Config{
		Message: "Config 1",
	}

	return (*ConfigA)(config)
}

func NewConfigAlternative() *ConfigB {
	config := &Config{
		Message: "Config 1",
	}
	return (*ConfigB)(config)
}

// AppConfig adalah struct yang menyimpan konfigurasi aplikasi
type AppConfig struct {
	AppName string
	Version int
	Port    int
	DBName  string
}

// PrintConfig mencetak konfigurasi aplikasi
func (a *AppConfig) PrintConfig() {
	fmt.Printf("App Name: %s, Version: %v\n", a.AppName, a.Version)
}
