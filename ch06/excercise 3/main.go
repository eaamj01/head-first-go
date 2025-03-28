// pcolumn prints the contents of a specified column from a
// CSV file. It takes two command-line arguments: the name
// of the file to read, and the column number to print.
// Example:
//
// pcolumn my.csv 2
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// printColumn reads the specified file and prints the
// specified column from each row. A non-nil return
// value indicates an error was encountered.
func printColumn(fileName string, column int) error {
	// Open the file for reading, and return any error.
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	// Ensure the file gets closed, even if there's an
	// error. We'll talk about "defer" in chapter 12.
	defer file.Close()
	// Set up a new csv.Reader that reads from the file.
	reader := csv.NewReader(file)
	// Loop until an error is encountered.
	for {
		// Read the next row.
		row, err := reader.Read()
		// Reaching the end of the file is an "error", but
		// it's an expected one, so just return.
		if err == io.EOF {
			return nil
		} else if err != nil {
			// Return any other type of error, because
			// those are NOT expected.
			return err
		}
		// If a column outside the row was requested,
		// return an error.
		if int(column) > len(row) {
			return fmt.Errorf("invalid column: %d", column)
		}
		// Otherwise, print the requested column.
		fmt.Println(row[column-1])
	}
}

// check logs an error and exits if the error is not nil.
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	arguments := os.Args[1:]
	if len(arguments) != 2 {
		log.Fatal("Incorrect number of arguments")
	}

	fileName := arguments[0]
	columnNum, err := strconv.ParseInt(arguments[1], 10, 64)
	check(err)
	err = printColumn(fileName, int(columnNum))
	check(err)
	// Call printColumn with the file name and column
	// number. Note that ParseInt returns an int64 and
	// printColumn wants an int, so you'll need to cast
	// the type.
	// Call "check" on the error value returned from
	// printColumn.
}
