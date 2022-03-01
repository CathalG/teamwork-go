package main

import "testing"

func TestCustomerImporter(t *testing.T){

    got := CustomerImporter("customers.csv")
    want := 10

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}