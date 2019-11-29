package main

import "fmt"

var ROMAN_SYMBOLS = map[byte]int {
	'I':1,
	'V':5,
	'X':10,
	'L':50,
	'C':100,
	'D':500,
	'M':1000,
}

func romanToInt(s string) (res int) {

	res+=ROMAN_SYMBOLS[s[0]]

	for i := 1; i < len(s);i++ {
		if ROMAN_SYMBOLS[s[i-1]] < ROMAN_SYMBOLS[s[i]] {
			res -= ROMAN_SYMBOLS[s[i-1]]
			res += ROMAN_SYMBOLS[s[i]] - ROMAN_SYMBOLS[s[i-1]]
		} else {
			res += ROMAN_SYMBOLS[s[i]]
		}
	}
	return
}

func main() {
	fmt.Printf("Roman %s to digit %d\n", "MCMXCIV", romanToInt("MCMXCIV"))
}


