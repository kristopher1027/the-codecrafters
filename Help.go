package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convtodec(num string, base int) (int64, error) {
	return strconv.ParseInt(num, base, 64)
}

func main() {
	fmt.Println("Base Converter (type 'quit' to exit)")
	
	for {
		var num, base string
		fmt.Print("> ")
		fmt.Scan(&num)
		
		if num == "quit" {
			break
		}
		
		fmt.Scan(&base)
		base = strings.ToLower(base)
		
		switch base {
		case "bin":
			if val, err := convtodec(num, 2); err == nil {
				fmt.Println("Decimal:", val)
			} else {
				fmt.Println("Error: '"+num+"' is not valid binary")
			}
			
		case "hex":
			if val, err := convtodec(num, 16); err == nil {
				fmt.Println("Decimal:", val)
			} else {
				fmt.Println("Error: '"+num+"' is not valid hex")
			}
			
		case "dec":
			if val, err := convtodec(num, 10); err == nil {
				if val < 0 {
					val = -val
					fmt.Println("Binary: -" + strconv.FormatInt(val, 2))
					fmt.Println("Hex: -" + strings.ToUpper(strconv.FormatInt(val, 16)))
				} else {
					fmt.Println("Binary:", strconv.FormatInt(val, 2))
					fmt.Println("Hex:", strings.ToUpper(strconv.FormatInt(val, 16)))
				}
			} else {
				fmt.Println("Error: '"+num+"' is not valid decimal")
			}
			
		default:
			fmt.Println("Error: Use bin, hex, or dec")
		}
	}
}
