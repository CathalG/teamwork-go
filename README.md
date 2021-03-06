# Teamwork-go

A code assessment submission for Teamwork. Details of the backend code challenge below:

*package customerimporter reads from the given customers.csv file and returns a
sorted (data structure of your choice) of email domains along with the number
of customers with e-mail addresses for each domain.  Any errors should be
logged (or handled). Performance matters (this is only ~3k lines, but *could*
be 1m lines or run on a small machine).*

## Installation & Instructions for use

**Clone the project:**
```
git clone git@github.com:CathalG/teamwork-go.git
```

**Run the project:**
Assuming Go is installed on your machine:
```
go run interview.go
```

**Expected Output:**
The main function of the application will output the results of the CSV customer import to the command line in the format shown below:
```
Domain: php.net => Users: 2
Domain: census.gov => Users: 7 
Domain: digg.com => Users: 2   
Domain: meetup.com => Users: 4
```
## Testing
Unit and performance tests can be found in:
```
interview_test.go
```
Running Tests: I've included a table-driven test to validate several expected outputs, these can be found/modified in 'countTests'.
To run:
```
go test
```
There's also a benchmark test to evaluate the performance of customerimporter:
```
go test -bench=.
```
Sample benchmark output:
```
$ go test -bench=.
goos: windows
goarch: amd64
pkg: teamwork-go-tests
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkCustomerImporter-8          470           2399312 ns/op
PASS
ok      teamwork-go-tests       1.577s
```


