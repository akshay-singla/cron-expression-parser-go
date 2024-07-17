package corn

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// Define max values for each field
	maxMinute     int = 59
	maxHour       int = 23
	maxDayOfMonth int = 31
	maxMonth      int = 12
	maxDayOfWeek  int = 7
)

// FieldParser interface for parsing fields
type FieldParser interface {
	Parse(min, max int) ([]string, error)
}

// WildcardParser for '*'
type WildcardParser struct{}

func (p WildcardParser) Parse(min, max int) ([]string, error) {
	var values []string
	for i := min; i <= max; i++ {
		values = append(values, strconv.Itoa(i))
	}
	return values, nil
}

// StepParser for fields with '/'
type StepParser struct {
	Base string
	Step int
}

func (p StepParser) Parse(min, max int) ([]string, error) {
	var values []string
	step := p.Step
	if p.Base == "*" {
		for i := min; i <= max; i += step {
			values = append(values, strconv.Itoa(i))
		}
		return values, nil
	}
	baseValues, err := getFieldParser(p.Base).Parse(min, max)
	if err != nil {
		return nil, err
	}
	for i, v := range baseValues {
		if i%step == 0 {
			values = append(values, v)
		}
	}
	return values, nil
}

// ListParser for fields with ','
type ListParser struct {
	List []string
}

func (p ListParser) Parse(min, max int) ([]string, error) {
	var values []string
	for _, v := range p.List {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		if num < min || num > max {
			return nil, fmt.Errorf("value %d out of range [%d, %d]", num, min, max)
		}
		values = append(values, strconv.Itoa(num))
	}
	return values, nil
}

// RangeParser for fields with '-'
type RangeParser struct {
	Start, End int
}

func (p RangeParser) Parse(min, max int) ([]string, error) {
	var values []string
	if p.Start < min || p.End > max {
		return nil, fmt.Errorf("range %d-%d out of range [%d, %d]", p.Start, p.End, min, max)
	}
	for i := p.Start; i <= p.End; i++ {
		values = append(values, strconv.Itoa(i))
	}
	return values, nil
}

// SingleValueParser for single values
type SingleValueParser struct {
	Value int
}

func (p SingleValueParser) Parse(min, max int) ([]string, error) {
	if p.Value < min || p.Value > max {
		return nil, fmt.Errorf("value %d out of range [%d, %d]", p.Value, min, max)
	}
	return []string{strconv.Itoa(p.Value)}, nil
}

// Factory function to get the appropriate parser
func getFieldParser(field string) FieldParser {
	if field == "*" {
		return WildcardParser{}
	}
	if strings.Contains(field, "/") {
		parts := strings.Split(field, "/")
		step, _ := strconv.Atoi(parts[1])
		return StepParser{Base: parts[0], Step: step}
	}
	if strings.Contains(field, ",") {
		list := strings.Split(field, ",")
		return ListParser{List: list}
	}
	if strings.Contains(field, "-") {
		rangeParts := strings.Split(field, "-")
		start, _ := strconv.Atoi(rangeParts[0])
		end, _ := strconv.Atoi(rangeParts[1])
		return RangeParser{Start: start, End: end}
	}
	num, _ := strconv.Atoi(field)
	return SingleValueParser{Value: num}
}

type Parser struct {
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
		return Parser{}, fmt.Errorf("invalid number of fields: expected 6, got %d", len(fields))
	}

	minute := fields[0]
	hour := fields[1]
	dayOfMonth := fields[2]
	month := fields[3]
	dayOfWeek := fields[4]
	command := fields[5]

	minuteExpanded, err := getFieldParser(minute).Parse(0, maxMinute)
	if err != nil {
		return Parser{}, fmt.Errorf("error parsing minute field: %v", err)
	}
	hourExpanded, err := getFieldParser(hour).Parse(0, maxHour)
	if err != nil {
		return Parser{}, fmt.Errorf("error parsing hour field: %v", err)
	}
	dayOfMonthExpanded, err := getFieldParser(dayOfMonth).Parse(1, maxDayOfMonth)
	if err != nil {
		return Parser{}, fmt.Errorf("error parsing day of month field: %v", err)
	}
	monthExpanded, err := getFieldParser(month).Parse(1, maxMonth)
	if err != nil {
		return Parser{}, fmt.Errorf("error parsing month field: %v", err)
	}
	dayOfWeekExpanded, err := getFieldParser(dayOfWeek).Parse(0, maxDayOfWeek)
	if err != nil {
		return Parser{}, fmt.Errorf("error parsing day of week field: %v", err)
	}

	return Parser{
		Minute:     minuteExpanded,
		Hour:       hourExpanded,
		DayOfMonth: dayOfMonthExpanded,
		Month:      monthExpanded,
		DayOfWeek:  dayOfWeekExpanded,
		Command:    command,
	}, nil
}
