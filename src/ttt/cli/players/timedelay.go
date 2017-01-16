package players

import "time"

type TimeDelay struct{}

func (timeDelay TimeDelay) delay(duration int) {
	time.Sleep(duration * time.Second)
}
