package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/akshay-singla/cron-expression-parser-go/corn"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: cron-parser \"*/15 0 1,15 * 1-5 /usr/bin/find\"")
		return
	}

	cronExpr := os.Args[1]
	fields := strings.Fields(cronExpr)

	if len(fields) != 6 {
		fmt.Println("Invalid cron expression format.")
		return
	}

	parser, err := corn.Validate(fields)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%-14s%s\n", "minute", strings.Join(parser.Minute, " "))
	fmt.Printf("%-14s%s\n", "hour", strings.Join(parser.Hour, " "))
	fmt.Printf("%-14s%s\n", "day of month", strings.Join(parser.DayOfMonth, " "))
	fmt.Printf("%-14s%s\n", "month", strings.Join(parser.Month, " "))
	fmt.Printf("%-14s%s\n", "day of week", strings.Join(parser.DayOfWeek, " "))
	fmt.Printf("%-14s%s\n", "command", parser.Command)

}
