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

	parser.Print()

}
