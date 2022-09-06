package utility

import (
	"strings"
	"time"
)

func IsToday(date string) bool {
	today := time.Now()
	return strings.Compare(today.Format("2006-01-02"), date) == 0
}
