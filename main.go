package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Flags struct {
	inputStr       string
	inputBinaryStr string
}

func (f *Flags) Define() {
	flag.StringVar(&f.inputStr, "s", "", "enter input for string to 8-bit binary conversion")
	flag.StringVar(&f.inputBinaryStr, "b", "", "enter input for to 8-bit binary to string conversion")
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

func binaryToString(bytes string) (string, error) {
	b := make([]byte, 0)
	for _, field := range strings.Fields(bytes) {
		v, err := strconv.ParseUint(field, 2, 8)
		if err != nil {
			return "", err
		}
		b = append(b, byte(v))
	}
	return string(b), nil
}

func main() {
	f := Flags{}
	f.Define()
	f.Parse()
	var input, output string
	if output = strings.TrimSpace(f.inputStr); output != "" {
		input = f.inputStr
		output = stringToBinary(f.inputStr)
	} else if output = strings.TrimSpace(f.inputBinaryStr); output != "" {
		input = f.inputBinaryStr
		output, _ = binaryToString(f.inputBinaryStr)
	}
	fmt.Printf("Input: %s\nOutput: %s\n", input, output)
}
