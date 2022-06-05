package main

import "time"

type TestFactory struct{}

var testFactory TestFactory

func (tf *TestFactory) NewSingleVestingRecord() *VestingRecord {
	t, _ := time.Parse(layout, "2020-03-04")
	return &VestingRecord{
		Action:          "VEST",
		EmployeeID:      "E001",
		EmployeeName:    "Alice Smith",
		VestingID:       "ISO-001",
		Date:            t,
		VestingQuantity: float64(10),
	}
}

func (tf *TestFactory) NewTwoVestingRecord() []*VestingRecord {
	t, _ := time.Parse(layout, "2020-04-04")
	res := make([]*VestingRecord, 0)
	res = append(res, tf.NewSingleVestingRecord())
	res = append(res, &VestingRecord{
		Action:          "VEST",
		EmployeeID:      "E001",
		EmployeeName:    "Alice Smith",
		VestingID:       "ISO-001",
		Date:            t,
		VestingQuantity: float64(20),
	})
	return res
}

func (tf *TestFactory) NewCancelVestingRecord() *VestingRecord {
	t, _ := time.Parse(layout, "2020-03-04")
	return &VestingRecord{
		Action:          "Cancel",
		EmployeeID:      "E001",
		EmployeeName:    "Alice Smith",
		VestingID:       "ISO-001",
		Date:            t,
		VestingQuantity: float64(40),
	}
}
