package main

import "time"

type TestFactory struct{}

var testFactory TestFactory

func (tf *TestFactory) NewVestingRecordForSinglePerson() *VestingRecord {
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

func (tf *TestFactory) NewVestingRecordForSinglePersonDifferentID() *VestingRecord {
	t, _ := time.Parse(layout, "2020-04-04")
	return &VestingRecord{
		Action:          "VEST",
		EmployeeID:      "E001",
		EmployeeName:    "Alice Smith",
		VestingID:       "ISO-002",
		Date:            t,
		VestingQuantity: float64(50),
	}
}

func (tf *TestFactory) NewMultiVestingRecordsForSinglePerson() []*VestingRecord {
	t, _ := time.Parse(layout, "2020-04-04")
	res := make([]*VestingRecord, 0)
	res = append(res, tf.NewVestingRecordForSinglePerson())
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

func (tf *TestFactory) NewVestingRecordForMultiPeople() []*VestingRecord {
	t, _ := time.Parse(layout, "2021-04-04")
	res := make([]*VestingRecord, 0)
	res = append(res, tf.NewVestingRecordForSinglePerson())
	res = append(res, &VestingRecord{
		Action:          "VEST",
		EmployeeID:      "E002",
		EmployeeName:    "Bobby Jones",
		VestingID:       "NSO-001",
		Date:            t,
		VestingQuantity: float64(20),
	})
	return res
}

func (tf *TestFactory) NewMultiVestingRecordsForMultiPeople() []*VestingRecord {
	t, _ := time.Parse(layout, "2021-04-04")
	res := make([]*VestingRecord, 0)
	res = append(res, tf.NewMultiVestingRecordsForSinglePerson()...)
	res = append(res, &VestingRecord{
		Action:          "VEST",
		EmployeeID:      "E002",
		EmployeeName:    "Bobby Jones",
		VestingID:       "NSO-001",
		Date:            t,
		VestingQuantity: float64(20),
	})
	return res
}

func (tf *TestFactory) NewCancelVestingRecord() *VestingRecord {
	t, _ := time.Parse(layout, "2020-04-04")
	return &VestingRecord{
		Action:          "CANCEL",
		EmployeeID:      "E001",
		EmployeeName:    "Alice Smith",
		VestingID:       "ISO-001",
		Date:            t,
		VestingQuantity: float64(40),
	}
}
