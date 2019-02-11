package main

import (
	"flag"
	"fmt"
)

type Flags struct {
	input string
}

func (f *Flags) Define() {
	flag.StringVar(&f.input, "s", "", "enter input for string to 8-bit binary conversion")
}

func (f *Flags) Parse() {
	flag.Parse()
}

func stringToBinary(s string) (b string) {
	for _, c := range s {
		b = fmt.Sprintf("%s %.8b", b, c)
	}
	return
}

func main() {
	f := Flags{}
	f.Define()
	f.Parse()
	b := stringToBinary(f.input)
	fmt.Printf("Original: %s\nBinary: %s\n", f.input, b)
}
