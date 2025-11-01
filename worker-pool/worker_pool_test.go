package workerpool

import "testing"

func TestBasicWorkerPool(t *testing.T) {
	BasicWorkerPool()
}

func TestProcessJobs(t *testing.T) {
	jobs := []int{1, 2, 3, 4, 5}

	results := ProcessJobs(3, jobs)

	if len(results) != 5 {
		t.Errorf("expected 5 results, got %d", len(results))
	}

	sum := 0

	for _, r := range results {
		sum += r
	}

	expectedSum := (1 + 2 + 3 + 4 + 5) * 2

	if sum != expectedSum {
		t.Errorf("expected sum %d, got %d", expectedSum, sum) // 결과 합계 확인 (순서는 보장 안 됨)
	}
}
