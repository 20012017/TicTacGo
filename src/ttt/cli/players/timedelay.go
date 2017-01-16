package players

import "time"

type TimeDelay struct{}

func (timeDelay TimeDelay) delay(duration int) {
	time.Sleep(time.Duration(duration) * time.Second)
}
