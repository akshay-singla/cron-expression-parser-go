package corn

import (
	"reflect"
	"strconv"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		input    []string
		expected Parser
		err      bool
	}{
		{
			input: []string{"*", "*", "*", "*", "*", "/path/to/command"},
			expected: parser{
				Minute:     generateSequence(0, maxMinute),
				Hour:       generateSequence(0, maxHour),
				DayOfMonth: generateSequence(1, maxDayOfMonth),
				Month:      generateSequence(1, maxMonth),
				DayOfWeek:  generateSequence(0, maxDayOfWeek),
				Command:    "/path/to/command",
			},
			err: false,
		},
		{
			input: []string{"*/15", "0", "1,15", "1-5", "*", "/path/to/command"},
			expected: parser{
				Minute:     []string{"0", "15", "30", "45"},
				Hour:       []string{"0"},
				DayOfMonth: []string{"1", "15"},
				Month:      []string{"1", "2", "3", "4", "5"},
				DayOfWeek:  generateSequence(0, maxDayOfWeek),
				Command:    "/path/to/command",
			},
			err: false,
		},
		{
			input: []string{"0", "12", "10-20", "2", "1-5", "/path/to/command"},
			expected: parser{
				Minute:     []string{"0"},
				Hour:       []string{"12"},
				DayOfMonth: []string{"10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"},
				Month:      []string{"2"},
				DayOfWeek:  []string{"1", "2", "3", "4", "5"},
				Command:    "/path/to/command",
			},
			err: false,
		},
		{
			input:    []string{"60", "*", "*", "*", "*", "/path/to/command"},
			expected: nil,
			err:      true,
		},
		{
			input:    []string{"*", "*", "0", "*", "*", "/path/to/command"},
			expected: nil,
			err:      true,
		},
		{
			input:    []string{"*", "*", "*", "*", "8", "/path/to/command"},
			expected: nil,
			err:      true,
		},
		{
			input: []string{"*/2", "*", "*", "*", "0-7", "/path/to/command"},
			expected: parser{
				Minute:     generateStepSequence(0, maxMinute, 2),
				Hour:       generateSequence(0, maxHour),
				DayOfMonth: generateSequence(1, maxDayOfMonth),
				Month:      generateSequence(1, maxMonth),
				DayOfWeek:  generateSequence(0, maxDayOfWeek),
				Command:    "/path/to/command",
			},
			err: false,
		},
	}

	for _, test := range tests {
		result, err := Validate(test.input)
		if (err != nil) != test.err {
			t.Errorf("Validate(%v) returned error: %v, expected error: %v", test.input, err, test.err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Validate(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func generateSequence(start, end int) []string {
	var result []string
	for i := start; i <= end; i++ {
		result = append(result, strconv.Itoa(i))
	}
	return result
}

func generateStepSequence(start, end, step int) []string {
	var result []string
	for i := start; i <= end; i += step {
		result = append(result, strconv.Itoa(i))
	}
	return result
}
