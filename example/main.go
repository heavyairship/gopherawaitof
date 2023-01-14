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
