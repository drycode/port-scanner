package utils

import (
	"fmt"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"
)

var assertionStatement = "%s -- expected: %v | actual: %v"

// Equals... fails the test if exp is not equal to act.
func AssertEquals(tb testing.TB, name string, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		fmt.Printf(assertionStatement, name, exp, act)
		tb.FailNow()
	}
}

type SafeSlice struct {
	mu        sync.RWMutex
	OpenPorts []string
}

func (ss *SafeSlice) Length() int {
	ss.mu.RLock()
	defer ss.mu.RUnlock()
	return len(ss.OpenPorts)
}

func (ss *SafeSlice) Append(val string, wg *sync.WaitGroup) {
	ss.mu.Lock()
	ss.OpenPorts = append((*ss).OpenPorts, val)
	ss.mu.Unlock()
	wg.Done()
}

// GetUlimit ...
func GetUlimit() (int, error) {
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		return -1, err
	}
	ulimit := strings.TrimSpace(string(out))
	i, err := strconv.ParseInt(ulimit, 10, 64)
	if err != nil {
		return -1, err
	}
	return int(i), nil
}
