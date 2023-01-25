package server

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/caarlos0/env"
)

const (
	defaultAddress       string        = "127.0.0.1:53000"
	defaultDBAddress     string        = "postgres://localhost/mydb?sslmode=disable"
	defaultTokenLifeTime time.Duration = time.Minute * 2
	defaultConfig        string        = "config.json"
)

type Config struct {
	Address       string        `env:"ADDRESS"`
	DBAddress     string        `env:"DB_ADDRESS"`
	ConfigFile    string        `env:"CONFIG"`
	TokenLifeTime time.Duration `env:"TOKEN_DURATION"`
	Environment   string        `env:"ENVIRONMENT"`
}

type ConfigFile struct {
	Address       string        `json:"address"`
	DBAddress     string        `json:"db_address"`
	TokenLifeTime time.Duration `json:"token_duration"`
}

func (config *ConfigFile) UnmarshalJSON(b []byte) error {
	type MyTypeAlias ConfigFile

	unmarshalledJSON := &struct {
		*MyTypeAlias
		TokenLifeTime string `json:"token_duration"`
	}{
		MyTypeAlias: (*MyTypeAlias)(config),
	}
	err := json.Unmarshal(b, &unmarshalledJSON)
	if err != nil {
		return err
	}

	config.TokenLifeTime, err = time.ParseDuration(unmarshalledJSON.TokenLifeTime)
	if err != nil {
		return err
	}

	return nil
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

	if c.TokenLifeTime == defaultTokenLifeTime && cfgFromFile.TokenLifeTime != 0 {
		c.TokenLifeTime = cfgFromFile.TokenLifeTime
	}

	return nil
}

func GetConfig() (*Config, error) {
	c := &Config{}

	flag.StringVar(&c.Address, "a", defaultAddress, "Socket to listen on")
	flag.StringVar(&c.DBAddress, "d", defaultDBAddress, "Database address")
	flag.DurationVar(&c.TokenLifeTime, "t", defaultTokenLifeTime, "User token lifetime duration")
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
