package date

import . "time"

func GetDate(date int64) Time {
	// Convert int to unix date
	tm := Unix(date, 0)
	return tm
}
