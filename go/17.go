/*
If the numbers 1 to 5 are written out in words: one, two, three, four, five,
then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in
words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and
forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20
letters. The use of "and" when writing out numbers is in compliance with British
usage.

Answer: 21124
*/

package main

import "fmt"

var num2eng = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twleve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func english(n int) string {
	if n >= 100 {
		hundreds := num2eng[n/100] + "hundred"
		tens := english(n - n/100*100)
		if len(tens) > 0 {
			return hundreds + "and" + tens
		} else {
			return hundreds
		}
	}
	if value, ok := num2eng[n]; ok {
		return value
	}

	if n >= 20 {
		return num2eng[n/10*10] + num2eng[n-n/10*10]
	}

	return ""
}

func e17() {
	s := "onethousand"
	for n := 1; n < 1000; n++ {
		s += english(n)
	}
	fmt.Println(len(s))
}

func init() {
    register("17", e17)
}
