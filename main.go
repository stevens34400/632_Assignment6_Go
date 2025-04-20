package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// Task represents a unit of work
type Task struct {
	Name string
}

// process simulates computation
func (t Task) process() string {
	time.Sleep(500 * time.Millisecond)
	return fmt.Sprintf("%s processed by %s", t.Name, getGoroutineID())
}

// getGoroutineID returns a string identifier (not actual goroutine ID, just for demonstration)
func getGoroutineID() string {
	return fmt.Sprintf("goroutine-%d", time.Now().UnixNano()%10000)
}

func main() {
	tasks := make(chan Task, 20)
	var wg sync.WaitGroup
	resultFile := "output/results.txt"

	// Create result output directory if not exists
	os.MkdirAll("output", os.ModePerm)
	file, err := os.Create(resultFile)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer file.Close()
	mutex := &sync.Mutex{} // to synchronize writes to file

	// Add tasks to channel
	for i := 1; i <= 20; i++ {
		tasks <- Task{Name: fmt.Sprintf("Task-%d", i)}
	}
	close(tasks) // close after sending all tasks

	// Start workers
	workerCount := 4
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for task := range tasks {
				result := task.process()
				mutex.Lock()
				_, err := file.WriteString(result + "\n")
				mutex.Unlock()
				if err != nil {
					log.Printf("Worker-%d error writing result: %v", id, err)
				} else {
					log.Printf("Worker-%d completed: %s", id, result)
				}
			}
		}(i + 1)
	}

	wg.Wait()
	log.Println("All tasks completed.")
}