package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	period := s[len(s)-2:]
	s = s[:len(s)-2]

	parts := strings.Split(s, ":")
	hourStr, min, sec := parts[0], parts[1], parts[2]

	hour, _ := strconv.Atoi(hourStr)

	if period == "AM" {
		if hour == 12 {
			hour = 0
		}
	} else {
		if hour != 12 {
			hour += 12
		}
	}

	hourStr = fmt.Sprintf("%02d", hour)

	return fmt.Sprintf("%s:%s:%s", hourStr, min, sec)
}

func main() {
	time := "12:01:00AM"
	result := timeConversion(time)
	fmt.Println(result)
}
