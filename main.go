package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

type VestingRecord struct {
	Action          string
	EmployeeID      string
	EmployeeName    string
	VestingID       string
	Date            time.Time
	VestingQuantity float64
}

type VestingRecordSummary struct {
	employeeID   string
	employeeName string
	vestingID    string
	totalVested  float64
}

func (vrs *VestingRecordSummary) EmployeeID() string {
	return vrs.employeeID
}

func (vrs *VestingRecordSummary) EmployeeName() string {
	return vrs.employeeName
}

func (vrs *VestingRecordSummary) VestingID() string {
	return vrs.vestingID
}

func (vrs *VestingRecordSummary) TotalVested() float64 {
	return vrs.totalVested
}

func (vrs *VestingRecordSummary) Vesting(optionNum float64) {
	vrs.totalVested += optionNum
}

func (vrs *VestingRecordSummary) CancelVesting(optionNum float64) {
	vrs.totalVested -= optionNum
	if vrs.totalVested < 0 {
		panic(fmt.Sprintf("Number of cancel exceed number vested for employee ID %v", vrs.employeeID))
	}
}

func NewVestingRecordSummary(employeeID, employeeName, vestingID string) *VestingRecordSummary {
	return &VestingRecordSummary{
		employeeID:   employeeID,
		employeeName: employeeName,
		vestingID:    vestingID,
	}
}

type InputReader interface {
	Next() (*VestingRecord, error)
}

type OutputWriter interface {
	PrintRecord(record map[string]*VestingRecordSummary)
}

type VestingReport struct {
	inputReader   InputReader
	outputWriter  OutputWriter
	vestingRecord map[string]*VestingRecordSummary
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <csvFile>\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}
	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println(`Please run the command with the following input: path to the csv file, search date in format "YYYY-MM-DD", and the optional precision value between 0 and 6`)
		os.Exit(1)
	}
	inputReader, err := NewCSVInputReader(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	targetDate, err := time.Parse(layout, os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	precision := 0
	if len(os.Args) == 4 {
		precision, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if precision < 0 || precision > 6 {
			fmt.Println("The precision need to be between 0 and 6")
			os.Exit(1)
		}
	}
	outputWriter := NewSTDOUTWriter(precision)
	vestingReport := NewVestingReport(inputReader, outputWriter)
	err = vestingReport.RunReport(targetDate)
	if err != nil {
		fmt.Println(err)
	}
}
