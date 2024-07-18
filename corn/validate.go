package corn

import (
	"fmt"
	"strings"
)

type Parser interface {
	Print()
}

type parser struct {
	Minute     []string
	Hour       []string
	DayOfMonth []string
	Month      []string
	DayOfWeek  []string
	Command    string
}

// Validate function to parse the fields and return a Parser struct
func Validate(fields []string) (Parser, error) {
	if len(fields) != 6 {
		return nil, fmt.Errorf("invalid number of fields: expected 6, got %d", len(fields))
	}

	minute := fields[0]
	hour := fields[1]
	dayOfMonth := fields[2]
	month := fields[3]
	dayOfWeek := fields[4]
	command := fields[5]

	minuteExpanded, err := getFieldParser(minute).Parse(0, maxMinute)
	if err != nil {
		return nil, fmt.Errorf("error parsing minute field: %v", err)
	}
	hourExpanded, err := getFieldParser(hour).Parse(0, maxHour)
	if err != nil {
		return nil, fmt.Errorf("error parsing hour field: %v", err)
	}
	dayOfMonthExpanded, err := getFieldParser(dayOfMonth).Parse(1, maxDayOfMonth)
	if err != nil {
		return nil, fmt.Errorf("error parsing day of month field: %v", err)
	}
	monthExpanded, err := getFieldParser(month).Parse(1, maxMonth)
	if err != nil {
		return nil, fmt.Errorf("error parsing month field: %v", err)
	}
	dayOfWeekExpanded, err := getFieldParser(dayOfWeek).Parse(0, maxDayOfWeek)
	if err != nil {
		return nil, fmt.Errorf("error parsing day of week field: %v", err)
	}

	return parser{
		Minute:     minuteExpanded,
		Hour:       hourExpanded,
		DayOfMonth: dayOfMonthExpanded,
		Month:      monthExpanded,
		DayOfWeek:  dayOfWeekExpanded,
		Command:    command,
	}, nil
}

func (p parser) Print() {
	fmt.Printf("%-14s%s\n", "minute", strings.Join(p.Minute, " "))
	fmt.Printf("%-14s%s\n", "hour", strings.Join(p.Hour, " "))
	fmt.Printf("%-14s%s\n", "day of month", strings.Join(p.DayOfMonth, " "))
	fmt.Printf("%-14s%s\n", "month", strings.Join(p.Month, " "))
	fmt.Printf("%-14s%s\n", "day of week", strings.Join(p.DayOfWeek, " "))
	fmt.Printf("%-14s%s\n", "command", p.Command)
}
