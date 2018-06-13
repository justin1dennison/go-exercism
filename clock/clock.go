package clock

import "fmt"

type Clock struct {
	h int
	m int
}

func New(hours, minutes int) Clock {
	if minutes < 0 {
		hours = hours + (minutes / 60) - 1
		minutes = 60 + (minutes % 60)
	}
	if hours < 0 {
		hours = (hours % 24) + 24
	}

	seconds := hours*3600 + minutes*60
	h := seconds / 3600
	m := (seconds - h*3600) / 60
	return Clock{h % 24, m}

}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) Add(min int) Clock {
	seconds := c.h*3600 + (c.m+min)*60
	hours := (seconds / 3600)
	minutes := (seconds - 3600*hours) / 60
	c.h = hours % 24
	c.m = minutes
	return c

}

func (c Clock) Subtract(min int) Clock {
	seconds := c.h*3600 + (c.m-min)*60
	hours := seconds / 3600
	minutes := (seconds - 3600*hours) / 60
	if minutes < 0 {
		hours += (minutes / 60) - 1
		minutes = (minutes % 60) + 60
	}
	if hours < 0 {
		hours = (hours % 24) + 24
	}
	c.h = hours % 24
	c.m = minutes

	return c
}
