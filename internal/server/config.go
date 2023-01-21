package server

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env"
)

const (
	defaultAddress   string = "127.0.0.1:53000"
	defaultDBAddress string = "postgres://localhost/mydb?sslmode=disable"
	defaultConfig    string = "config.json"
)

type Config struct {
	Address    string `env:"ADDRESS"`
	DBAddress  string `env:"DB_ADDRESS"`
	ConfigFile string `env:"CONFIG"`
}

type ConfigFile struct {
	Address   string `json:"address"`
	DBAddress string `json:"db_address"`
}

func loadConfigFromFile(c *Config) error {
	if c.ConfigFile == "" {
		return nil
	}

	log.Printf("Loading config from file '%s'", c.ConfigFile)
	fileBytes, err := os.ReadFile(c.ConfigFile)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var cfgFromFile ConfigFile
	err = json.Unmarshal(fileBytes, &cfgFromFile)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if c.Address == defaultAddress && cfgFromFile.Address != "" {
		c.Address = cfgFromFile.Address
	}

	if c.DBAddress == defaultDBAddress && cfgFromFile.DBAddress != "" {
		c.DBAddress = cfgFromFile.DBAddress
	}

	return nil
}

func GetConfig() (*Config, error) {
	c := &Config{}

	flag.StringVar(&c.Address, "a", defaultAddress, "Socket to listen on")
	flag.StringVar(&c.DBAddress, "d", defaultDBAddress, "Database address")
	flag.StringVar(&c.ConfigFile, "config", defaultConfig, "Config file name")
	flag.StringVar(&c.ConfigFile, "c", defaultConfig, "Config file name")
	flag.Parse()

	err := env.Parse(c)
	if err != nil {
		log.Fatal(err)
	}

	err = loadConfigFromFile(c)
	if err != nil {
		log.Fatal(err)
	}

	return c, nil
}
