package config

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	DbHost          string
	DbName          string
	DbUser          string
	DbPassword      string
	DbPort          string
	StanClusterName string
}

func Get() *Config {
	cfg := &Config{
		DbHost:          "localhost",
		DbName:          "L0",
		DbPort:          "5432",
		DbUser:          "user",
		DbPassword:      "pass",
		StanClusterName: "test-cluster",
	}

	jBytes, err := json.MarshalIndent(cfg, "", "   ")
	if err != nil {
		panic(err)
	}

	fmt.Println("Config:\n" + string((jBytes)))

	return cfg
}
