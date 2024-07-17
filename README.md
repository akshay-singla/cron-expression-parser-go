# Cron Expression Parser

This is a command-line application written in Go that parses a cron string and expands each field to show the times at which it will run.

## Installation

### Prerequisites

- Go programming language installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

### Clone the Repository

Clone this repository to your local machine using Git:

```bash
git clone https://github.com/akshay-singla/cron-expression-parser-go.git
cd cron-expression-parser-go
```

### Build the Application
Build the Go application using the following command:

```bash
go build
```

### Usage
Run the application with a cron expression as a single argument:

```bash
./cron-expression-parser-go "*/15 0 1,15 * 1-5 /usr/bin/find"
```

Replace **"*/15 0 1,15 * 1-5 /usr/bin/find"** with your desired cron expression.

The output will be formatted as a table with each field name followed by the times it will run:

```bash
minute         0 15 30 45
hour           0
day of month   1 15
month          1 2 3 4 5 6 7 8 9 10 11 12
day of week    1 2 3 4 5
command       /usr/bin/find
```

### Cron Expression Format
The cron string should adhere to the standard cron format with five time fields (minute, hour, day of month, month, day of week) plus a command.

- minute: Allowed values: 0-59, */X (every X minutes), X,Y (specific minutes)
- hour: Allowed values: 0-23, */X (every X hours), X,Y (specific hours)
- day of month: Allowed values: 1-31, */X (every X days), X,Y (specific days)
- month: Allowed values: 1-12, */X (every X months), X,Y (specific months)
- day of week: Allowed values: 0-7 (both 0 and 7 represent Sunday)
