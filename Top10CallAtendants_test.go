package main

import (
	mdl "AVAYA/models"
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	var att mdl.Attendant
	var atts []mdl.Attendant
	att.Name = "1"
	att.TotalMonthCalls = 10
	atts = append(atts, att)
	att = mdl.Attendant{}
	att.Name = "2"
	att.TotalMonthCalls = 5
	atts = append(atts, att)
	res, _ := Sort(atts)
	fmt.Print(res)
}
