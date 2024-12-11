package main

import (
    "fmt"
    "log"
    
    "npkg.dev/encoding/yaml.v3"
)

// Define struct that matches your YAML structure
type Config struct {
    Server struct {
        Host string `yaml:"host"`
        Port int    `yaml:"port"`
    } `yaml:"server"`
    Database struct {
        Username string `yaml:"username"`
        Password string `yaml:"password"`
        DBName   string `yaml:"dbname"`
    } `yaml:"database"`
}

func main() {
    // Example YAML data
    yamlData := []byte(`
      server:
        host: localhost
        port: 8080
      database:
        username: admin
        password: secret123
        dbname: myapp
    `)

    // Create a Config struct to hold the parsed data
    var config Config

    // Parse YAML data
    err := yaml.Unmarshal(yamlData, &config)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    // Access the parsed data
    fmt.Printf("Server Host: %s\n", config.Server.Host)
    fmt.Printf("Server Port: %d\n", config.Server.Port)
    fmt.Printf("Database Username: %s\n", config.Database.Username)
    fmt.Printf("Database Name: %s\n", config.Database.DBName)

    // Marshal back to YAML
    data, err := yaml.Marshal(&config)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    fmt.Printf("\nMarshaled YAML:\n%s", string(data))
}
