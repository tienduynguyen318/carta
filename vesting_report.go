package main

import (
	"fmt"
	"time"
)

func NewVestingReport(inputReader InputReader, outputWriter OutputWriter) *VestingReport {
	return &VestingReport{
		inputReader:   inputReader,
		outputWriter:  outputWriter,
		vestingRecord: make(map[string]VestingRecordSummary),
	}
}

func (vr *VestingReport) RunReport(targetDate time.Time) error {
	var err error
	for vr.inputReader.HasNext() {
		vestingRecord, err := vr.inputReader.Next()
		if err != nil {
			return err
		}
		key := fmt.Sprintf("%v-%v", vestingRecord.EmployeeID, vestingRecord.VestingID)
		if _, ok := vr.vestingRecord[key]; !ok {
			initialQuantity := 0
			vr.vestingRecord[key] = VestingRecordSummary{
				EmployeeID:   vestingRecord.EmployeeID,
				EmployeeName: vestingRecord.EmployeeName,
				VestingID:    vestingRecord.VestingID,
				TotalVested:  &initialQuantity,
			}
		}
		switch vestingRecord.Action {
		case "VEST":
			if vestingRecord.Date.Unix() <= targetDate.Unix() {
				summary := vr.vestingRecord[key]
				*summary.TotalVested += vestingRecord.VestingQuantity

			}
		case "CANCEL":
			if vestingRecord.Date.Unix() <= targetDate.Unix() {
				summary := vr.vestingRecord[key]
				*summary.TotalVested -= vestingRecord.VestingQuantity
				if *summary.TotalVested < 0 {
					err = fmt.Errorf("Number of share vested smaller than cancellation quantity")
					return err
				}
			}
		}
	}
	for _, summary := range vr.vestingRecord {
		vr.outputWriter.Print(summary)
	}
	return err
}
