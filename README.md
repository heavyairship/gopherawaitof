# gopherawaitof

Provides a generic function `ForAwaitOf[T any](tasks []func() T, handler func(t T))` that mimics the behavior of the [`for await..of`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/for-await...of) construct in JavaScript.

# Example

Here's an example.

```
package main

import (
	"fmt"
	"time"

	"github.com/gopherawaitof/gopherawaitof"
)

func main() {
	start := time.Now()
	tasks := []func() string{
		func() string {
			time.Sleep(time.Second * 3)
			return fmt.Sprintf("task 1 finished at t=%f", time.Since(start).Seconds())
		}, func() string {
			time.Sleep(time.Second * 2)
			return fmt.Sprintf("task 2 finished at t=%f", time.Since(start).Seconds())
		}, func() string {
			time.Sleep(time.Second * 1)
			return fmt.Sprintf("task 3 finished at t=%f", time.Since(start).Seconds())
		}, func() string {
			time.Sleep(time.Second * 5)
			return fmt.Sprintf("task 4 finished at t=%f", time.Since(start).Seconds())
		},
	}
	handler := func(result string) { fmt.Printf(result+", handled at t=%f\n", time.Since(start).Seconds()) }
	gopherawaitof.ForAwaitOf(tasks, handler)
	fmt.Printf("executed all tasks in %f seconds\n", time.Since(start).Seconds())
}
```

Let's run the program.

```
λ go run example/main.go 
task 1 finished at t=3.001314, handled at t=3.001405
task 2 finished at t=2.004487, handled at t=3.001443
task 3 finished at t=1.003386, handled at t=3.001447
task 4 finished at t=5.004719, handled at t=5.004800
executed all tasks in 5.004836 seconds
```

In the above example, the four tasks are executed concurrently, but the results are handled sequentially in task order as they become available. Crucially, result handling can begin before all tasks are complete.

Formally, suppose we have an ordered list of tasks `task_i` for `i` in `1..N`, a set of results `result_j` for `j` in `1..N` where `result_j` is the result for `task_i` iff `i = j`, and a handler function `handler`. Then, `gopherawaitof.ForAwaitOf` guarantees the following:
1. `handler(result_m)` happens-before `handler(result_n)` iff `m < n` for all `m, n` in `1..N`.
2. `handler(result_i)` happens-before `task_i` execution for all `i` in `1..N`.
