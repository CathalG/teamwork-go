package main

import(
    "testing"
    "strconv"
)

type CountTest struct {
    domain string; expected int
}

var countTests = []CountTest{
    CountTest{"github.io", 8},
    CountTest{"mac.com", 7},
    CountTest{"bandcamp.com", 3},
    CountTest{"huffingtonpost.com", 10},
    CountTest{"google.ca", 7},
    CountTest{"unc.edu", 8},
    CountTest{"yelp.com", 4},
    CountTest{"microsoft.com", 7},
    CountTest{"bbb.org", 5},
}

//Testing domain counts
func TestCustomerImporter(t *testing.T) {
    for _, test := range countTests {
        if output := customerImporter("customers.csv")[test.domain] ; output != test.expected {
            t.Errorf("Count %q not equal to expected %q", strconv.Itoa(output), strconv.Itoa(test.expected))
        }
    }
}

// Benchmark performance test
func BenchmarkCustomerImporter(b *testing.B){
    for i :=0; i < b.N ; i++{
        customerImporter("customers.csv")
    }
}