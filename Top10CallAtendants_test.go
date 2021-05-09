package main

import (
	mdl "AVAYA/models"
	"fmt"
	"regexp"
	"testing"
)

func TestSort(t *testing.T) {
	var att mdl.Attendant
	var atts []mdl.Attendant
	att.Name = "1"
	att.TotalMonthCalls = 20
	atts = append(atts, att)
	att = mdl.Attendant{}
	att.Name = "2"
	att.TotalMonthCalls = 8
	atts = append(atts, att)

	want := regexp.MustCompile("\\**(20 calls){1}\\.*")
	fmt.Print(want)
	res, err := Sort(atts)
	if !want.MatchString(res) || err != nil {
		t.Fatal("Erro", res, err, want)
	}
}
