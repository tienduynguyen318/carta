package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type inputReader struct {
	data    [][]string
	counter int
}

func NewInputReader(filePath string) (*inputReader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return &inputReader{data: data}, nil
}

var layout = "2006-01-02"

func (ir *inputReader) Next() (VestingRecord, error) {
	var err error
	record := ir.data[ir.counter]
	ir.counter += 1
	if len(record) != 6 {
		err = fmt.Errorf("Malformed data")
		return VestingRecord{}, err
	}
	quantity, err := strconv.Atoi(record[5])
	if err != nil {
		return VestingRecord{}, err
	}
	date, err := time.Parse(layout, record[4])
	if err != nil {
		return VestingRecord{}, err
	}
	return VestingRecord{
		Action:          record[0],
		EmployeeID:      record[1],
		EmployeeName:    record[2],
		VestingID:       record[3],
		Date:            date,
		VestingQuantity: quantity,
	}, err
}

func (ir *inputReader) HasNext() bool {
	if ir.counter >= len(ir.data) {
		return false
	}
	return true
}
