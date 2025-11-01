package main

import (
	"runtime"
	"sync/atomic"
)

type Mutex interface {
	Lock()
	Unlock()
	Trylock() bool
}

// 스핀락 기반 뮤텍스
// 스핀락은 락을 얻을 때까지 계속 반복해서 시도하는 락. (busy waiting)
// 구현이 간단하고 락 대기 시간이 짧으면 효율적이지만, CPU를 계속 사용하는 단점이 존재한다.
type SpinMutex struct {
	state int32 //0:unlocked , 1; locked
}

func (m *SpinMutex) Lock() {
	//CAS 연산을 이용해서 락 획득
	for { //성공할 때까지 반복 (블로킹 동작)
		if atomic.CompareAndSwapInt32(&m.state, 0, 1) {
			//state가 unlock이면 lock으로 변경하고 락 획득 성공
			return
		}
		//락 획득 실패시 CPU 양보
		runtime.Gosched()
	}
}

func (m *SpinMutex) Unlock() {
	//state를 unlock으로 변경
	atomic.StoreInt32(&m.state, 0)
}

func (m *SpinMutex) Trylock() bool {
	//논블라킹 락 시도
	return atomic.CompareAndSwapInt32(&m.state, 0, 1)
}
