package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	mdl "AVAYA/models"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Loads attendants from an excel file extracting name and setting ids.
func LoadAttendants(qtt int) []mdl.Attendant {

	f, err := excelize.OpenFile("saopaulo.xlsx")
	if err != nil {
		println(err.Error())
		return nil
	}
	var attendant mdl.Attendant
	var attendants []mdl.Attendant
	id := 0
	for _, name := range f.GetSheetMap() {
		cell := f.GetCellValue(name, "B8")
		txt := fmt.Sprint(cell)
		index := 2
		for txt != "" && id < qtt {
			cell := f.GetCellValue("Sheet1", "B"+strconv.Itoa(index))
			txt = fmt.Sprint(cell)
			//println(id, txt)
			index += 23
			if !strings.Contains(txt, "SERVIDOR PUBLICO MUNICIPAL") {
				id++
				attendant.Id = id
				attendant.Name = txt
				attendants = append(attendants, attendant)
			}
		}
	}
	return attendants
}

// Loads calls for each attendant.
// Assertion 1: 20 days per month - from 1 to 31
// Assertion 2: 8 hours per day
// Assertion 3: A call lasts from 5 to 45 minutes.
func LoadCalls(attendants []mdl.Attendant) []mdl.Attendant {
	var daysOfWork map[int][]mdl.Call
	var calls []mdl.Call
	var call mdl.Call
	var loadedAttendants []mdl.Attendant
	for _, attendant := range attendants {
		qttDays := 1
		daysOfWork = make(map[int][]mdl.Call)
		totalMonthCalls := 0
		for qttDays <= 20 {
			day := rand.Intn(30) + 1
			if _, ok := daysOfWork[day]; !ok {
				duration := 480
				for duration > 0 {
					call.Day = day
					call.Duration = rand.Intn(40) + 5
					call.Attendant = attendant
					calls = append(calls, call)
					duration -= call.Duration
					totalMonthCalls++
				}
				daysOfWork[day] = calls
				qttDays++
			}
		}
		attendant.DaysOfWork = daysOfWork
		attendant.TotalMonthCalls = totalMonthCalls
		loadedAttendants = append(loadedAttendants, attendant)
	}
	return loadedAttendants
}
