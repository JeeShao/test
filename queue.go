/**
* @author Jee
* @date 2021/2/4 22:57
 */
package main

import (
	"fmt"

	"k8s.io/client-go/util/workqueue"
)

type Task struct {
	queue workqueue.RateLimitingInterface
}

func main() {
	task := &Task{workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "task")}
	defer task.queue.ShutDown()

	task.queue.Add(1)
	task.queue.Add(2)
	task.queue.Add(3)

	fmt.Println(task.queue.Len()) //3

	e, shutdown := task.queue.Get()
	task.queue.Add(1)
	fmt.Println(task.queue.Len()) //2

	task.queue.Done(e)
	fmt.Println(task.queue.Len()) //3

	fmt.Println(e, shutdown)
}
