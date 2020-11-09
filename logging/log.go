package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)


func init() {
    logrus.SetOutput(os.Stdout)
}