package utils

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestSetPercentage(t *testing.T) {
	ss := SafeSlice{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go ss.append(strconv.Itoa(i), &wg)
	}
	wg.Wait()

	if ss.length() != 1000 {
		fmt.Print(len(ss.OpenPorts))
		t.Fail()
	}

}
