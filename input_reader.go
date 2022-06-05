package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type CSVInputReader struct {
	cvsReader *csv.Reader
}

func NewCSVInputReader(filePath string) (*CSVInputReader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	return &CSVInputReader{cvsReader: reader}, nil
}

var layout = "2006-01-02"

func (cir *CSVInputReader) Next() (*VestingRecord, error) {
	var err error
	record, err := cir.cvsReader.Read()
	if err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}
	if len(record) != 6 {
		err = fmt.Errorf("Malformed data")
		return nil, err
	}
	quantity, err := strconv.ParseFloat(record[5], 64)
	if err != nil {
		return nil, err
	}
	date, err := time.Parse(layout, record[4])
	if err != nil {
		return nil, err
	}
	return &VestingRecord{
		Action:          record[0],
		EmployeeID:      record[1],
		EmployeeName:    record[2],
		VestingID:       record[3],
		Date:            date,
		VestingQuantity: quantity,
	}, nil
}
