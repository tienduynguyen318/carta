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
	inputReader, err := NewCSVInputReader(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	precision, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(err)
	}
	outputWriter := NewSTDOUTWriter(precision)
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
