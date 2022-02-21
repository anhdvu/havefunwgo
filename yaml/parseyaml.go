package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name      string
	Cookie    []string
	TestCard  TestCard
	Share     SharedConfig
	TestCases []TestCase
}

type TestCard struct {
	Number     string
	ExpiryDate string
	Cvv        string
	Pin        string
}

type SharedConfig struct {
	AmountMin      int    `yaml:"amountMin"`
	AmountMax      int    `yaml:"amountMax"`
	DefaultCode    string `yaml:"defaultCode"`
	DefaultDecimal string `yaml:"defaultDecimal"`
}

type TestCase struct {
	Included                      bool
	Name                          string
	Runs                          int
	Mode                          string
	ATM                           bool
	SettleType                    string
	Reversal                      string
	Mcc                           string
	Source                        string
	Foreign                       bool
	OriginalCurrencyCode          string
	OriginalCurrencyDecimalPlaces string
	Acquirer                      string
	Province                      string
	Country                       string
}

func (c *Config) ParseConfig() error {
	f, err := os.Open("config.yaml")
	if err != nil {
		return err
	}
	defer f.Close()
	d := yaml.NewDecoder(f)
	err = d.Decode(c)
	if err != nil {
		return err
	}
	return nil
}

func ParseConfig() (*Config, error) {
	raw, err := os.ReadFile("config.yaml")

	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(raw, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func main() {
	t := time.Now()
	c := &Config{}
	c.ParseConfig()
	fmt.Println(time.Since(t))
	fmt.Printf("%+v\n", c)

	t = time.Now()
	c, _ = ParseConfig()
	fmt.Println(time.Since(t))
}
