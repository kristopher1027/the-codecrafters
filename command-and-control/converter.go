package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HexToDec(hexstr string) int64 {
	hex, error := strconv.ParseInt(hexstr, 16, 64)
	if error != nil {
		fmt.Println("error message")
	}
	return hex
}
func BinToDec(binstr string) int64 {
	bin, error := strconv.ParseInt(binstr, 2, 64)
	if error != nil {
		fmt.Println("invalid base")
	}
	return bin
}
func FromDec(decc string) (string, string) {
	dec, err := strconv.ParseInt(decc, 10, 64)
	if err != nil {
		fmt.Println("invalid input")
	}
	deci := strconv.FormatInt(dec, 2)
	hex := strconv.FormatInt(dec, 16)
	return deci, strings.ToUpper(hex)
}
