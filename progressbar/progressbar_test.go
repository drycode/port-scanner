package progressbar

import (
	"fmt"
	"testing"
)


func TestSetPercentage(t *testing.T) {
	var tests = []struct {
        total, scanned, expected int
    }{
        {100, 10, 10},
        {200, 20, 10},
		{2345, 233, 9},
		{16500000, 1, 0},
	}
	for _, param := range tests {
		testname := fmt.Sprintf("%d, %d", param.total, param.scanned)
		t.Run(testname, func(t *testing.T) {
			value := calculatePercentage(param.scanned, param.total)
			if value != param.expected{
				fmt.Print(value)
				t.Fail()
			}
		})
	}
	

}