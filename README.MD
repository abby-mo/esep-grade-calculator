# ESEP Grade Calculator

A  Go library that computes final course grades. Grades come in three types (assignments, exams, essays). The library averages each kind, applies a weight per type, and converts the weighted result to a letter grade.

## Repository layout

```
.
├── README.MD
└── grade-calculator/          <- Go module; run all go commands from here
    ├── go.mod                 <- module name (esep/grade-calculator) and Go version
    ├── grade_calculator.go    <- library code: GradeCalculator, Grade, GradeType
    └── grade_calculator_test.go <- unit tests 
```

## Runtime requirements

- Go 1.24 or newer (`go version` to check). 
- No third-party dependencies

## Usage

```go
gc := NewGradeCalculator()
gc.AddGrade("hw 1", 95, Assignment)
gc.AddGrade("midterm", 88, Exam)
gc.AddGrade("essay 1", 91, Essay)
fmt.Println(gc.GetFinalGrade()) // letter grade, ex. "A"
```

## Run the unit tests

```
cd grade-calculator/
go test ./...
```

Add `-v` to see each test by name, and `-run TestName` to run a single test.

## Run the unit tests with code coverage

Print the coverage percentage:

```
cd grade-calculator/
go test -cover ./...
```

Write a coverage profile, see per-function coverage, and open a line-by-line
HTML report (covered lines green, uncovered lines red):

```
cd grade-calculator/
go test -coverprofile=coverage.txt ./...
go tool cover -func=coverage.txt
go tool cover -html=coverage.txt
```

`coverage.txt` is a generated file and should not be committed.
