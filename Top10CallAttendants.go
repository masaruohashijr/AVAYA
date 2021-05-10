package main

import (
	mdl "AVAYA/models"
	u "AVAYA/utils"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Provides sorting functions of the attendants collections by total of call per month.
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

// Calls the sort function used and formats the output in string.
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

// Applies the merge sort algorithm to split a bigger array in halves, successively, until the length reachs 1,
// so the merge function is called in sequence, positioning each attendant in crescent order of total call per month.
func mergeSort(attendants []mdl.Attendant, channelMS chan []mdl.Attendant) {
	if len(attendants) == 1 {
		channelMS <- attendants
		return
	}
	channelLeft := make(chan []mdl.Attendant)
	channelRight := make(chan []mdl.Attendant)
	// Splits the attendants array into two parts: left and right.
	left, right := attendants[0:len(attendants)/2], attendants[len(attendants)/2:]
	// goroutine mergeSort
	go mergeSort(left, channelLeft)
	go mergeSort(right, channelRight)
	left, right = <-channelLeft, <-channelRight
	// the sender is the only sender of the channel, so we dont need to close the channel politely or gracefully.
	close(channelLeft)
	close(channelRight)
	mergeChan := make(chan []mdl.Attendant)
	go merge(left, right, mergeChan)
	channelMS <- <-mergeChan
	close(mergeChan)
	return
}

// Assemble the merged array correctly positioning each attendant from the halves being gone through.
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
		// TotalMonthCalls
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
