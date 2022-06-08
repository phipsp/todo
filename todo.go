package todo

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorCyan  = "\033[36m"

	dateFormat = "2006/01/02"
)

type Logger interface {
	Printf(format string, v ...any)
}

var logger Logger

func init() {
	logger = log.New(os.Stdout, "[TODO] ", 0)
}

func SetLogger(l Logger) {
	logger = l
}

// UntilDate alarms the user if the passed in date has passed. The format of the date has to be in "2006/01/02".
// The alarm will be reported in two ways:
// - logged to the packages default logger which can be overwritten by the user.
// - returned as an error to be handled manually by the user.
func UntilDate(date, description string) error {
	endDate, err := time.Parse(dateFormat, date)
	if err != nil {
		return err
	}

	if !time.Now().Before(endDate) {
		msg := fmt.Sprintf("%sThis should have been done by now!%s\n%s\t%s%s\n", colorRed, colorReset, colorCyan, description, colorReset)
		logger.Printf(msg)
		return errors.New(msg)
	}
	return nil
}
