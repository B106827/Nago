package utils

import (
    "time"
)

// 現在日時からの差分（日単位）を整形し返す（diff:-1,timeFormat:mだと1ヶ月前)
func GetAddDateTime(diff int, timeFormat string) time.Time {
    now := time.Now()
    diffY, diffM, diffD := 0, 0, 0
    switch timeFormat {
        case "y":
            diffY, diffM, diffD = diff, 0, 0
        case "m":
            diffY, diffM, diffD = 0, diff, 0
        case "d":
            diffY, diffM, diffD = 0, 0, diff
    }
    return now.AddDate(diffY, diffM, diffD)
}
