package tests

import (
	"strconv"
	"sync"
	"testing"

	. "github.com/drypycode/portscanner/utils"
)

func TestSafeSlice(t *testing.T) {
	ss := SafeSlice{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go ss.Append(strconv.Itoa(i), &wg)
	}
	wg.Wait()
	AssertEquals(t, "Check all ports evaluated", ss.Length(), 1000)
}
