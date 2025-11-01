package workerpool

import "fmt"

func BasicWorkerPool() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//워커 3개 시작
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//작업 9개 보내기
	for j := 1; j <= 9; j++ {
		jobs <- j
	}

	close(jobs) //더 이상 작업 없음을 알림

	//결과 9개 받기
	for a := 1; a <= 9; a++ {
		result := <-results
		fmt.Println("결과 : ", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("워커 %d: 작업 %d 시작\n", id, j)
		result := j * 2
		fmt.Printf("워커 %d: 작업 %d 완료\n", id, j)
		results <- result
	}
}
