/*
You are given the following information, but you may prefer to do some research
for yourself.

    1 Jan 1900 was a Monday.
    Thirty days has September,
    April, June and November.
    All the rest have thirty-one,
    Saving February alone,
    Which has twenty-eight, rain or shine.
    And on leap years, twenty-nine.
	A leap year occurs on any year evenly divisible by 4, but not on a century
		unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1
Jan 1901 to 31 Dec 2000)?

Answer: 171
*/

package main

import "fmt"

func isLeap(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	}

	return year%4 == 0
}

func numDays(month, year int) int {
	switch month {
	case 2:
		if isLeap(year) {
			return 29
		}
		return 28
	case 4, 6, 9, 11:
		return 30
	}
	return 31
}

func e19() {
	weekday := 1 // Monday
	day, month, year := 1, 1, 1900
	count := 0

	for {
		if year > 2000 {
			break
		}

		if year > 1900 {
			if weekday == 0 && day == 1 {
				count += 1
			}
		}

		day += 1
		weekday = (weekday + 1) % 7
		if day > numDays(month, year) {
			day = 1
			month += 1
			if month > 12 {
				month = 1
				year += 1
			}
		}
	}

	fmt.Println(count)
}

func init() {
    register("19", e19)
}
