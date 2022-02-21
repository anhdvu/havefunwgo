package main

import "fmt"

type Payload interface {
	method() string
}

type Auth struct {
	Method string
	Params struct {
		Province string
		Country  string
		Mcc      string
	}
}

func (a *Auth) method() string {
	return a.Method
}

type Settle struct {
	Method string
	Params struct {
		Province string
		Country  string
		Mcc      string
	}
}

func (s *Settle) method() string {
	return s.Method
}

type Defaults struct {
	DefaultProvince string
	DefaultCountry  string
	DefaultMcc      string
}

type Config struct {
	Province string
	Country  string
	Mcc      string
}

func main() {
	a := &Auth{}
	d := Defaults{"HANOI", "VNM", "7299"}
	c := Config{"Conway", "", ""}

	handleFieldsWithDefaults(a, d, c)
	fmt.Printf("%+v", a)
}

func handleFieldsWithDefaults(p Payload, d Defaults, c Config) {
	switch t := p.(type) {
	case *Auth:
		if c.Province != "" {
			t.Params.Province = c.Province
		} else {
			t.Params.Province = d.DefaultProvince
		}

		if c.Country != "" {
			t.Params.Country = c.Country
		} else {
			t.Params.Country = d.DefaultCountry
		}

		if c.Mcc != "" {
			t.Params.Mcc = c.Mcc
		} else {
			t.Params.Mcc = d.DefaultMcc
		}
	}
}
