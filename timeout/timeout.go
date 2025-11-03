package timeout

import "time"

func TimeoutWithDuration(workDuration, timeout time.Duration) string {
	result := make(chan string)

	go func() {
		time.Sleep(workDuration)
		result <- "done"
	}()

	select {
	case res := <-result:
		return res
	case <-time.After(timeout):
		return "timeout"
	}
}
