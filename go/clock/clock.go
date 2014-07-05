package clock

import (
	"fmt"
)

type Clock struct {
	minutes int
}

const (
	MINUTES_PER_DAY = 60 * 24
)

func (c Clock) String() string {
	hours := c.minutes / 60
	remainding_minutes := c.minutes % 60
	return fmt.Sprintf("%02d:%02d", hours, remainding_minutes)
}

func (c Clock) Add(minutes int) Clock {
	return Clock{
		minutes: normalize_minutes((c.minutes + minutes) % MINUTES_PER_DAY),
	}
}

func New(hours, minutes int) Clock {
	return Clock{
		minutes: hours*60 + minutes,
	}
}

func normalize_minutes(minutes int) int {
	if minutes >= 0 {
		return minutes
	}
	return MINUTES_PER_DAY + minutes
}
