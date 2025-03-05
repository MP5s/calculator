package application

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MP5s/calculator/pkg/dir"
)

// Структура конфигурации
type Config struct {
	Debug bool `json:"debug"`
	Web   bool `json:"web"`
}

// Создание новой конфигурации
func LoadConfig() *Config {
	configFilePath := dir.Json_file()
	fmt.Println(configFilePath)

	config := &Config{}
	file, err := os.Open(configFilePath)
	if err != nil {
		panic("Unable to open config file")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		panic(fmt.Sprintf("Failed to decode config file: %v", err))
	}
	return config
}
