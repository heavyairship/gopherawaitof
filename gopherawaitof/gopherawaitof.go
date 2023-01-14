package gopherawaitof

// Executes tasks concurrently, but handles results sequentially in task order.
// Result handling may begin before all tasks are complete.
func ForAwaitOf[T any](tasks []func() T, handler func(t T)) {
	queue := make(chan struct {
		int
		t T
	}, len(tasks))
	for i := 0; i < len(tasks); i++ {
		i := i
		go func() {
			queue <- struct {
				int
				t T
			}{i, tasks[i]()}
		}()
	}
	next := 0
	results := make(map[int]T, len(tasks))
	for i := 0; i < len(tasks); i++ {
		result := <-queue
		results[result.int] = result.t
		for {
			if t, ok := results[next]; ok {
				handler(t)
				next += 1
			} else {
				break
			}
		}
	}
}
