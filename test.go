package main

import (
	"fmt"
)

type Student struct {
	name string
}

func (s Student) calAvg(data []float64) (avgResult float64) {
	sum := 0.0
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	avgResult = sum / float64(len(data))
	return
}

func (s Student) judge(avg float64) (judgeResult string) {
	if avg >= 80 {
		judgeResult = "Excellent"
	} else {
		judgeResult = "Bad"
	}
	return
}

func main() {
	a001 := Student{"taka"}
	data := []float64{70, 80, 85, 90, 0}
	var avg float64 = a001.calAvg(data)
	result := a001.judge(avg)
	fmt.Println(avg)
	fmt.Println(a001.name, result)
}
