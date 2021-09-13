package datetime

import (
	"fmt"
	"time"
)

func CombineDateTime(d, t string) (time.Time, error) {
	completeStringFmt := "%sT%s+07:00"
	completeStringFmt = fmt.Sprintf(completeStringFmt, d, t)
	return time.Parse(time.RFC3339, completeStringFmt)
}
