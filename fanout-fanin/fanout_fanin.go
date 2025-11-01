package fanoutfanin

import "fmt"

// 팬 아웃: 하나의 입력 채널을 여러 워커가 읽음
func BasicFanOut() {
	input := make(chan int, 10)

	for w := 1; w <= 3; w++ {
		go func(id int) {
			for num := range input {
				fmt.Printf(" worker :%d 처리 : %d", id, num)
			}
		}(w)
	}

	//데이터 보내기
	for i := 1; i <= 9; i++ {
		input <- i
	}

	close(input)

	// 워커들이 끝날 때까지 대기
	// (실제로는 sync.WaitGroup 사용해야 함)
	fmt.Scanln()
}

// 팬 인 : 여러 채널을 하나로 합침
func BasicFanIn() {
	//입력 채널 3개
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 10; i <= 12; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	go func() {
		for i := 100; i <= 102; i++ {
			ch3 <- i
		}
		close(ch3)
	}()

	merged := fanIn(ch1, ch2, ch3)

	for v := range merged {
		fmt.Println("받음 ", v)
	}
}

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)

	//각 채널마다 고루틴 하나씩 생성
	for _, ch := range channels {
		go func(c <-chan int) {
			for v := range c {
				out <- v //하나의 채널로 합침
			}
		}(ch)
	}

	return out
}
