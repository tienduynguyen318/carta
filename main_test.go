package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type singleEventInputReaderDouble struct {
	counter       int
	resultCounter int
	record        *VestingRecord
}

func (rd *singleEventInputReaderDouble) Next() (*VestingRecord, error) {
	if rd.counter < rd.resultCounter {
		rd.counter++
		return rd.record, nil
	}
	return nil, nil
}

func NewSingleEventInputReaderDouble() *singleEventInputReaderDouble {
	singleRecord := testFactory.NewVestingRecordForSinglePerson()
	return &singleEventInputReaderDouble{
		resultCounter: 1,
		record:        singleRecord,
	}
}

type outputWriterDouble struct {
	builder   *strings.Builder
	precision int
}

func NewOutputWriterDouble(precision int) *outputWriterDouble {
	builder := new(strings.Builder)
	return &outputWriterDouble{builder: builder, precision: precision}
}

func (wd *outputWriterDouble) PrintRecord(records map[string]*VestingRecordSummary) {
	keys := make([]string, 0, len(records))
	for key := range records {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		summary := records[key]
		wd.builder.WriteString(fmt.Sprintf("%v,%v,%v,%v",
			summary.EmployeeID(),
			summary.EmployeeName(),
			summary.VestingID(),
			strconv.FormatFloat(summary.TotalVested(), 'f', wd.precision, 64),
		))
	}
}

func TestSinglePersonSuccessfulVest(t *testing.T) {
	inputDouble := NewSingleEventInputReaderDouble()
	outputDouble := NewOutputWriterDouble(0)
	vestingReport := NewVestingReport(inputDouble, outputDouble)
	targetDate, _ := time.Parse(layout, "2021-03-04")
	vestingReport.RunReport(targetDate)
	assert.Equal(t, "E001,Alice Smith,ISO-001,10", outputDouble.builder.String())
}

func TestSinglePersonFailVest(t *testing.T) {
	inputDouble := NewSingleEventInputReaderDouble()
	outputDouble := NewOutputWriterDouble(0)
	vestingReport := NewVestingReport(inputDouble, outputDouble)
	targetDate, _ := time.Parse(layout, "2019-03-04")
	vestingReport.RunReport(targetDate)
	assert.Equal(t, "E001,Alice Smith,ISO-001,0", outputDouble.builder.String())
}

type multiEventInputReaderDouble struct {
	counter       int
	resultCounter int
	records       []*VestingRecord
}

func (rd *multiEventInputReaderDouble) Next() (*VestingRecord, error) {
	if rd.counter < rd.resultCounter {
		res := rd.records[rd.counter]
		rd.counter++
		return res, nil
	}
	return nil, nil
}

func NewMultiEventInputReaderDouble(record []*VestingRecord) *multiEventInputReaderDouble {
	counter := len(record)
	return &multiEventInputReaderDouble{
		resultCounter: counter,
		records:       record,
	}
}

func TestSinglePersonMultiSuccessfulVestSameID(t *testing.T) {
	records := testFactory.NewMultiVestingRecordsForSinglePerson()
	inputDouble := NewMultiEventInputReaderDouble(records)
	outputDouble := NewOutputWriterDouble(0)
	vestingReport := NewVestingReport(inputDouble, outputDouble)
	targetDate, _ := time.Parse(layout, "2021-03-04")
	vestingReport.RunReport(targetDate)
	assert.Equal(t, "E001,Alice Smith,ISO-001,30", outputDouble.builder.String())
}

func TestSinglePersonMultiSuccessfulVestSameIDWithPrecision(t *testing.T) {
	records := make([]*VestingRecord, 0)
	record1 := testFactory.NewVestingRecordForSinglePerson()
	record1.VestingQuantity = 299.12332
	records = append(records, record1)
	record2 := testFactory.NewVestingRecordForSinglePerson()
	record2.VestingQuantity = 132.81241
	records = append(records, record2)
	inputDouble := NewMultiEventInputReaderDouble(records)
	outputDouble := NewOutputWriterDouble(2)
	vestingReport := NewVestingReport(inputDouble, outputDouble)
	targetDate, _ := time.Parse(layout, "2021-03-04")
	vestingReport.RunReport(targetDate)
	assert.Equal(t, "E001,Alice Smith,ISO-001,431.94", outputDouble.builder.String())
}

func TestSinglePersonMultiSuccessfulVestDifferentID(t *testing.T) {
	records := testFactory.NewMultiVestingRecordsForSinglePerson()
	records = append(records, testFactory.NewVestingRecordForSinglePersonDifferentID())
	inputDouble := NewMultiEventInputReaderDouble(records)
	outputDouble := NewOutputWriterDouble(0)
	vestingReport := NewVestingReport(inputDouble, outputDouble)
	targetDate, _ := time.Parse(layout, "2021-03-04")
	vestingReport.RunReport(targetDate)
	assert.Equal(t, "E001,Alice Smith,ISO-001,30E001,Alice Smith,ISO-002,50", outputDouble.builder.String())
}

func TestSinglePersonWithExceedCancelVest(t *testing.T) {
	records := make([]*VestingRecord, 0)
	records = append(records, testFactory.NewCancelVestingRecord())
	inputDouble := NewMultiEventInputReaderDouble(records)
	outputDouble := NewOutputWriterDouble(0)
	vestingReport := NewVestingReport(inputDouble, outputDouble)
	targetDate, _ := time.Parse(layout, "2021-03-04")
	assert.Panics(t, func() { vestingReport.RunReport(targetDate) })
}

func TestMultiPeopleSuccessfulVest(t *testing.T) {
	records := testFactory.NewMultiVestingRecordsForMultiPeople()
	inputDouble := NewMultiEventInputReaderDouble(records)
	outputDouble := NewOutputWriterDouble(0)
	vestingReport := NewVestingReport(inputDouble, outputDouble)
	targetDate, _ := time.Parse(layout, "2021-04-04")
	vestingReport.RunReport(targetDate)
	assert.Equal(t, "E001,Alice Smith,ISO-001,30E002,Bobby Jones,NSO-001,20", outputDouble.builder.String())
}
