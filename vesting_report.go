package main

import (
	"fmt"
	"time"
)

func NewVestingReport(inputReader InputReader, outputWriter OutputWriter) *VestingReport {
	return &VestingReport{
		inputReader:   inputReader,
		outputWriter:  outputWriter,
		vestingRecord: make(map[string]*VestingRecordSummary),
	}
}

func (vr *VestingReport) RunReport(targetDate time.Time) error {
	var err error
	for {
		vestingRecord, err := vr.inputReader.Next()
		if err != nil {
			return err
		}
		if vestingRecord == nil {
			break
		}
		key := fmt.Sprintf("%v-%v", vestingRecord.EmployeeID, vestingRecord.VestingID)
		if _, ok := vr.vestingRecord[key]; !ok {
			vr.vestingRecord[key] = NewVestingRecordSummary(vestingRecord.EmployeeID, vestingRecord.EmployeeName, vestingRecord.VestingID)
		}
		if vestingRecord.Date.Unix() <= targetDate.Unix() {
			switch vestingRecord.Action {
			case "VEST":
				vr.vestingRecord[key].Vesting(vestingRecord.VestingQuantity)
			case "CANCEL":
				vr.vestingRecord[key].CancelVesting(vestingRecord.VestingQuantity)
			}
		}
	}
	vr.outputWriter.PrintRecord(vr.vestingRecord)
	return err
}
