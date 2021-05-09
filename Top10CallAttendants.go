package main

import (
	mdl "AVAYA/models"
	u "AVAYA/utils"
	"fmt"
	"os"
	"strconv"
)

func main() {
	qttAttendants := 100
	if len(os.Args) >= 1 {
		qttAttendants, _ = strconv.Atoi(os.Args[1])
	}
	attendants := u.LoadAttendants(qttAttendants)
	attendants = u.LoadCalls(attendants)
	MS := make(chan []mdl.Attendant)
	go mergeSort(attendants, MS)
	r := <-MS
	fmt.Println("****************************")
	fmt.Println("*     TOP 10 Attendants    *")
	for i, v := range r {
		fmt.Printf("* %v \n", v)
		if i > 9 {
			break
		}
	}
	fmt.Println("****************************")

	close(MS)
}

func mergeSort(A []mdl.Attendant, MS chan []mdl.Attendant) {
	if len(A) == 1 {
		MS <- A
		return
	}
	leftChan := make(chan []mdl.Attendant)
	rightChan := make(chan []mdl.Attendant)
	left, right := A[0:len(A)/2], A[len(A)/2:]
	go mergeSort(left, leftChan)
	go mergeSort(right, rightChan)
	left, right = <-leftChan, <-rightChan
	close(leftChan)
	close(rightChan)
	mergeChan := make(chan []mdl.Attendant)
	go merge(left, right, mergeChan)
	MS <- <-mergeChan
	return
}

func merge(A, B []mdl.Attendant, MC chan []mdl.Attendant) (arr []mdl.Attendant) {
	arr = make([]mdl.Attendant, len(A)+len(B))
	j, k := 0, 0
	for i := 0; i < len(arr); i++ {
		if j >= len(A) {
			arr[i] = B[k]
			k++
			continue
		} else if k >= len(B) {
			arr[i] = A[j]
			j++
			continue
		}
		if A[j].TotalMonthCalls <= B[k].TotalMonthCalls {
			arr[i] = B[k]
			k++
		} else {
			arr[i] = A[j]
			j++
		}
	}
	MC <- arr
	return
}
