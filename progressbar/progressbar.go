package progressbar

import (
	"fmt"
	"strconv"
	"strings"
)

// ProgressBar ...
type ProgressBar struct {
	TotalPorts    	int
	Percentage    	int
	LastDisplayed 	int
	Output			[104]string
}

// NewProgressBar ...
func NewProgressBar(totalPorts int) ProgressBar {
	i:=1
	temp := [104]string{0: "[", 101: "]"}
	for i < 101 {
		temp[i] = "-"
		i ++
	}
	pb := ProgressBar{
		TotalPorts: totalPorts, 
		Percentage: 0, 
		LastDisplayed: 0,
		Output: temp,
	}
	
	return pb
}


// UpdatePercentage ...
func (pb *ProgressBar) UpdatePercentage(n int) {
	(*pb).Percentage = calculatePercentage(n, (*pb).TotalPorts)
	if (*pb).Percentage != (*pb).LastDisplayed {
		(*pb).renderView()
	}
}

func (pb *ProgressBar) renderView() {
	(*pb).LastDisplayed = (*pb).Percentage		
	(*pb).Output[102] = "  " + strconv.Itoa((*pb).LastDisplayed) + "%"
	(*pb).Output[(*pb).Percentage] = "#"	
	printable:=(*pb).Output[:]
	fmt.Print("\r" + strings.Join(printable, ""))
}

func calculatePercentage(scanned int, total int) int {	
	return int(float64(scanned) / float64(total) * 100)
}