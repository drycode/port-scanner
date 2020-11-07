package progressbar

import (
	"fmt"
	"testing"
)

func TestPercentageHelper(t *testing.T) {
	pb := ProgressBar{
		TotalPorts: 100, 
		Percentage: 0, 
		LastDisplayed: 0, 
		Output: [104]string{},
	}
	pb.PercentageHelper(60)
	if pb.Percentage != 60 {
		t.Fail()
	}

}
func TestSetPercentage(t *testing.T) {
	var tests = []struct {
        totalPorts, scanned, expected int
    }{
        {100, 10, 10},
        {200, 20, 10},
        {2345, 233, 9},
	}
	for _, param := range tests {
		testname := fmt.Sprintf("%d, %d", param.totalPorts, param.scanned)
		t.Run(testname, func(t *testing.T) {
			pb := ProgressBar{
				TotalPorts: param.totalPorts, 
				Percentage: 0, 
				LastDisplayed: 0, 
				Output: [104]string{},
			}
			pb.setPercentage(param.scanned)

			if pb.Percentage != param.expected {
				fmt.Print(pb.Percentage)
				t.Fail()
			}
		})
	}
	

}