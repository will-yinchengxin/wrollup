package wtools

import (
	"fmt"
	"strings"
	"time"
)

// ParseDuration 解析时间字符串，返回对应的时间戳
func ParseDuration(duration string) (int64, error) {
	duration = strings.ToLower(duration)
	if len(duration) < 2 {
		return 0, fmt.Errorf("invalid duration format: %s", duration)
	}

	value := duration[:len(duration)-1]
	unit := duration[len(duration)-1:]

	var num int
	_, err := fmt.Sscanf(value, "%d", &num)
	if err != nil {
		return 0, fmt.Errorf("invalid duration number: %s", value)
	}

	now := time.Now()
	var targetTime time.Time
	switch unit {
	case "m", "M":
		targetTime = now.AddDate(0, 0, 0-num*30) // 按30天计算一个月
	case "h", "H":
		targetTime = now.Add(-time.Duration(num) * time.Hour)
	case "d", "D":
		targetTime = now.AddDate(0, 0, -num)
	case "w", "W":
		targetTime = now.AddDate(0, 0, -num*7)
	case "y", "Y":
		targetTime = now.AddDate(-num, 0, 0)
	default:
		return 0, fmt.Errorf("unsupported duration unit: %s", unit)
	}

	return targetTime.UnixMilli(), nil
}
