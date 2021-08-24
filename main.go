package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"Goroutine/model"
	"Goroutine/utils"
)

func sendDataToChan(chanSend chan<- model.SimpleData, filePath string) {
	var (
		listData model.SimpleData
		wg       sync.WaitGroup
	)
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			data := utils.ReadFile(filePath)
			listData = utils.SplitString(data)
			chanSend <- listData
			wg.Done()
		}(&wg)
	}
	wg.Wait()
	close(chanSend)
}

func saveFile(done chan bool, chanReceive <-chan model.SimpleData) {
	var (
		count = 0
		wg    sync.WaitGroup
	)
	for content := range chanReceive {
		count++
		wg.Add(1)
		go utils.WriteFile("output"+strconv.Itoa(count)+".txt", content, &wg)
	}
	wg.Wait()
	done <- true
}

func run() {
	var (
		done     = make(chan bool)
		chanSend = make(chan model.SimpleData)
	)
	go saveFile(done, chanSend)
	go sendDataToChan(chanSend, "test.txt")
	<-done
}
func main() {
	start := time.Now()
	run()
	fmt.Println(time.Since(start))
}
