package main

import (
	"fmt"
	"strconv"
)

type STDOUTWriter struct {
	precision int
}

func (w *STDOUTWriter) PrintRecord(records map[string]*VestingRecordSummary) {
	for _, summary := range records {
		fmt.Printf("EmployeeID: %v, EmployeeName: %v, VestingID: %v, Quantity: %v. \n",
			summary.EmployeeID(),
			summary.EmployeeName(),
			summary.VestingID(),
			strconv.FormatFloat(summary.TotalVested(), 'f', w.precision, 64),
		)
	}
}

func NewSTDOUTWriter(precision int) *STDOUTWriter {
	return &STDOUTWriter{precision: precision}
}
