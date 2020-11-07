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
	temp := [104]string{0: "[", 102: "]"}
	for i < 102 {
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



// PercentageHelper ...
func (pb *ProgressBar) PercentageHelper(n int) {
	pb.setPercentage(n)
	if (*pb).Percentage != (*pb).LastDisplayed {
		(*pb).LastDisplayed = (*pb).Percentage		
		(*pb).Output[103] = "  " + strconv.Itoa((*pb).LastDisplayed) + "%"
		(*pb).Output[(*pb).Percentage + 1] = "#"	
		printable:=(*pb).Output[:]
		fmt.Print("\r" + strings.Join(printable, ""))
	}
}

func (pb *ProgressBar) setPercentage(i int) {	
	perc := int(float64(i) / float64((*pb).TotalPorts) * 100)
	(*pb).Percentage = perc
}