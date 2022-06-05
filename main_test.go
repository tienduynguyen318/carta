package main

import (
	"testing"
	"time"
)

type inputReaderDouble struct {
	singleRecord *VestingRecord
	twoRecord    []*VestingRecord
	cancelRecord *VestingRecord
}

func (rd *inputReaderDouble) Next() (*VestingRecord, error) {
	return rd.singleRecord, nil
}

func NewInputReaderDouble() *inputReaderDouble {
	singleRecord := testFactory.NewSingleVestingRecord()
	twoRecord := testFactory.NewTwoVestingRecord()
	cancelRecord := testFactory.NewCancelVestingRecord()
	return &inputReaderDouble{
		singleRecord: singleRecord,
		twoRecord:    twoRecord,
		cancelRecord: cancelRecord,
	}
}

type outputWriterDouble struct {
}

func NewOutputWriterDouble() *outputWriterDouble {
	return &outputWriterDouble{}
}

func (wd *outputWriterDouble) PrintRecord(record map[string]*VestingRecordSummary) {

}

func TestSingleEvent(t *testing.T) {
	t.Parallel()
	inputDouble := NewInputReaderDouble()
	outputWriter := NewOutputWriterDouble()
	vestingReport := NewVestingReport(inputDouble, outputWriter)
	targetDate, _ := time.Parse(layout, "2021-03-04")
	vestingReport.RunReport(targetDate)
}
