package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type VestingRecord struct {
	Action          string
	EmployeeID      string
	EmployeeName    string
	VestingID       string
	Date            time.Time
	VestingQuantity int
}

type VestingRecordSummary struct {
	EmployeeID   string
	EmployeeName string
	VestingID    string
	TotalVested  *int
}

type InputReader interface {
	Next() (VestingRecord, error)
	HasNext() bool
}

type OutputWriter interface {
	Print(interface{})
}

type VestingReport struct {
	inputReader   InputReader
	outputWriter  OutputWriter
	vestingRecord map[string]VestingRecordSummary
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <csvFile>\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}
	inputReader, err := NewInputReader(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	outputWriter, err := NewOutputWriter(os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
	vestingReport := NewVestingReport(inputReader, outputWriter)
	targetDate, err := time.Parse(layout, os.Args[2])
	if err != nil {
		fmt.Println(err)
	}
	err = vestingReport.RunReport(targetDate)
	if err != nil {
		fmt.Println(err)
	}
}
