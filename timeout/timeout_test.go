package timeout

import (
	"testing"
	"time"
)

func TestTimeoutWithDuration_Success(t *testing.T) {
	result := TimeoutWithDuration(500*time.Millisecond, 1*time.Second)

	if result != "done" {
		t.Errorf("expected '작업 완료', got '%s'", result)
	}
}

func TestTimeoutWithDuration_Timeout(t *testing.T) {
	// 작업 2초, 타임아웃 1초 → 타임아웃
	result := TimeoutWithDuration(2*time.Second, 1*time.Second)

	if result != "timeout" {
		t.Errorf("expected '타임아웃!', got '%s'", result)
	}
}
