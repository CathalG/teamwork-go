// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).

package main

import (
    "encoding/csv"
    "io"
    "log"
    "os"
	"regexp"
    "fmt"
)

func main(){
    count := customerImporter("customers.csv")
    fmt.Print(count)
}

func customerImporter (filepath string) map[string]int{
    // open file
    file, err := os.Open(filepath)
    
	//check for error in file read
	if err != nil {
        log.Fatal(err)
    }

    // close file on complete
    defer file.Close()

	//compile regex match before loop
	re := regexp.MustCompile(`@`)

	//declare domain count map
	domainCount := map[string]int{}

    // read CSV file
    csvReader := csv.NewReader(file)
    for {
        rec, err := csvReader.Read()

		// break loop if end of file
		if err == io.EOF {
            break
        }
		//if error string is set then log
        if err != nil {
            log.Fatal(err)
        }

        // splt the email string at @
		result := re.Split(rec[2], -1)

		//add result domainCount map, and add 1 to occurences
		for k, v := range result {
			 if k == 1 {
				domainCount[v] += 1; 
			 }
		}
    }
	// return domainCount
	return domainCount
}