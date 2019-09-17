package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }
func KToC(k Kelvin) Celsius     { return Celsius(k - 273.15) }

func CToK(c Celsius) Kelvin    { return Kelvin(c + 273.15) }
func FToK(f Fahrenheit) Kelvin { return Kelvin((f + 459.67) * 5 / 9) }

func (c Celsius) String() string { return fmt.Sprintf("%.3g°C", c) }

func (k Kelvin) String() string { return fmt.Sprintf("%.3g°K", k) }

type celsiusFlag struct{ Celsius }

type kelvinFlag struct{ Kelvin }

func (f *kelvinFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Kelvin = CToK(Celsius(value))
		return nil
	case "F", "°F":
		f.Kelvin = FToK(Fahrenheit(value))
		return nil
	case "K":
		f.Kelvin = Kelvin(value)
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func KelvinFlag(name string, value Kelvin, usage string) *Kelvin {
	f := kelvinFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Kelvin
}

var tempKelvin = KelvinFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*tempKelvin)
}
