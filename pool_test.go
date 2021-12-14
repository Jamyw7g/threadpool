package pool

import (
	"sync/atomic"
	"testing"
)

func TestPool(t *testing.T) {
	pool := NewPool(0)
	var num int32 = 0

	for i := 0; i < 32; i++ {
		pool.Add()
		go func() {
			defer pool.Done()
			atomic.AddInt32(&num, 1)
		}()
	}

	pool.Wait()

	if num != 32 {
		t.Errorf("Expect value 32, but result %d", num)
	}
}
