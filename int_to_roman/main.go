package main

import "fmt"

func main() {
	fmt.Println(intToRoman(3749))
}

func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	rom := ""
	for i := range vals {
		for num >= vals[i] {
			num -= vals[i]
			rom += syms[i]
		}
		i++
	}

	return rom
}
