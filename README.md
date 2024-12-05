# Advent of Code Runner

My solutions for AOC 2024 with basic CLI tool for running solutions with minimal boilerplate.
Solutions are mostly written fast and dirty only with the result in mind, so better not use them as reference ;)

## Usage

```bash
# Run all solutions
go run main.go

# Run specific day
go run main.go -day 5

# Run specific part
go run main.go -day 5 -part 1
```

## Adding New Solutions

1. Create a new package for your day:
```go
package dayX

type DayXSolver struct {
    input string
}

func NewDayXSolver(input string) Solver {
    return &DayXSolver{input: input}
}

func (d *DayXSolver) SolvePart1() int {
    // Your solution
}

func (d *DayXSolver) SolvePart2() int {
    // Your solution
}
```

2. Register in main.go:
```go
Register(X, func(input string) Solver { 
    return dayX.NewDayXSolver(input) 
})
```

## Structure
```
.
├── main.go           # Runner implementation
├── inputs/           # Input files
│   └── dX.txt       # Day X input
└── dayX/            # Day X solution
    └── solver.go
```