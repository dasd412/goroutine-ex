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

// 여러 작업 중 하나라도 먼저 완료되면 반환
func FirstResponse() string {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "server 1"
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "server 2"
	}()

	go func() {
		time.Sleep(250 * time.Millisecond)
		ch3 <- "server 3"
	}()

	select {
	case res := <-ch1:
		return res
	case res := <-ch2:
		return res
	case res := <-ch3:
		return res
	case <-time.After(1 * time.Second):
		return "all timeout"
	}
}
