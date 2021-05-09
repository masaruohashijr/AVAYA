package models

import "fmt"

type Attendant struct {
	Id              int
	Name            string
	DaysOfWork      map[int][]Call
	TotalMonthCalls int
}

type Call struct {
	Day       int
	Duration  int
	Attendant Attendant
}

func (c Call) String() string {
	return fmt.Sprintf("Attendant: %s - Day %d with %d minutes.", c.Attendant.Name, c.Day, c.Duration)
}

func (a Attendant) String() string {
	return fmt.Sprintf("Attendant: %s - %d calls.", a.Name, a.TotalMonthCalls)
}
