package progressbar

import (
	"fmt"
	"strconv"
	"strings"
)

// ProgressBar ...
type ProgressBar struct {
	TotalPorts    	int
	Current			int
	Percentage    	float64
	LastDisplayed 	string
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
		Current: 0, 
		Percentage: 0, 
		LastDisplayed: "",
		Output: temp,
	}
	
	return pb
}

func (p *ProgressBar) setPercentage(i int) {
	
	perc := float64(i) / float64((*p).TotalPorts)
	(*p).Percentage = perc
}

// PercentageHelper ...
func PercentageHelper(pb *ProgressBar, n int) {
	pb.setPercentage(n)
	
	currentPercentage := strconv.FormatInt(int64((*pb).Percentage*100), 10)
	if currentPercentage != (*pb).LastDisplayed {
		(*pb).LastDisplayed = strconv.FormatInt(int64((*pb).Percentage*100), 10)
		// fmt.Print("#")
		(*pb).Output[103] = "  " + (*pb).LastDisplayed + "%"
		(*pb).Output[int((*pb).Percentage * float64(100)) + 1] = "#"	
		printable:=(*pb).Output[:]
		fmt.Print("\r" + strings.Join(printable, ""))
	}

}