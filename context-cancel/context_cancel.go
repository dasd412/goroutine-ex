package contextcancel

import (
	"context"
	"fmt"
	"time"
)

func ContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("작업 취소 됨:", ctx.Err())
				return
			default:
				fmt.Println("working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

// 상위 작업 취소되면 하위 작업도 취소
func ParentChildCancel() []string {
	result := []string{}

	parentCtx, parentCancel := context.WithCancel(context.Background())
	defer parentCancel()

	childCtx1, _ := context.WithCancel(parentCtx)
	childCtx2, _ := context.WithCancel(parentCtx)

	go func() {
		<-childCtx1.Done()
		result = append(result, "child 1 canceled")
	}()

	go func() {
		<-childCtx2.Done()
		result = append(result, "child 2 canceled")
	}()

	time.Sleep(100 * time.Millisecond)
	parentCancel() // 부모 취소 되면 자식들도 자동 취소
	time.Sleep(100 * time.Millisecond)

	return result
}
