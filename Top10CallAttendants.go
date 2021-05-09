package main

import (
	mdl "AVAYA/models"
	u "AVAYA/utils"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	qttAttendants := 100
	var err error
	if len(os.Args) > 1 {
		qttAttendants, err = strconv.Atoi(os.Args[1])
		if err != nil {
			println("Error parsing to integer")
			qttAttendants = 100
		}
	}
	attendants := u.LoadAttendants(qttAttendants)
	attendants = u.LoadCalls(attendants)
	print(Sort(attendants))
}

func Sort(attendants []mdl.Attendant) (result string, err error) {
	channelMS := make(chan []mdl.Attendant)
	if len(attendants) == 0 {
		return "", errors.New("attendants is empty")
	}
	go mergeSort(attendants, channelMS)
	r := <-channelMS
	result = "****************************\n"
	result += fmt.Sprint("*     TOP 10 Attendants    *\n")
	for i, v := range r {
		result += fmt.Sprintf("* %v \n", v)
		if i > 9 {
			break
		}
	}
	result += fmt.Sprint("****************************\n")
	close(channelMS)
	return result, err
}

func mergeSort(attendants []mdl.Attendant, channelMS chan []mdl.Attendant) {
	if len(attendants) == 1 {
		channelMS <- attendants
		return
	}
	channelLeft := make(chan []mdl.Attendant)
	channelRight := make(chan []mdl.Attendant)
	left, right := attendants[0:len(attendants)/2], attendants[len(attendants)/2:]
	go mergeSort(left, channelLeft)
	go mergeSort(right, channelRight)
	left, right = <-channelLeft, <-channelRight
	close(channelLeft)
	close(channelRight)
	mergeChan := make(chan []mdl.Attendant)
	go merge(left, right, mergeChan)
	channelMS <- <-mergeChan
	close(mergeChan)
	return
}

func merge(left, right []mdl.Attendant, mergeChan chan []mdl.Attendant) (merged []mdl.Attendant) {
	merged = make([]mdl.Attendant, len(left)+len(right))
	j, k := 0, 0
	for i := 0; i < len(merged); i++ {
		if j >= len(left) {
			merged[i] = right[k]
			k++
			continue
		} else if k >= len(right) {
			merged[i] = left[j]
			j++
			continue
		}
		if left[j].TotalMonthCalls <= right[k].TotalMonthCalls {
			merged[i] = right[k]
			k++
		} else {
			merged[i] = left[j]
			j++
		}
	}
	mergeChan <- merged
	return
}
