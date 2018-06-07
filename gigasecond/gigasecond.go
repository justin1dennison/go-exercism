package gigasecond

// import path for the time package from the standard library
import "time"

func AddGigasecond(t time.Time) time.Time {
	seconds := time.Second * 1000000000
	return t.Add(seconds) 
}
