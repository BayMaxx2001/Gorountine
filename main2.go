package main

import (
	"sync"

	"Goroutine/model"
	"Goroutine/utils"
)

func sendDataToChan(chanSend chan model.SimpleData, filePath string) {

	data := utils.ReadFile(filePath)
	listData := utils.SplitString(data)

	chanSend <- listData
	close(chanSend)
}

func receiveDataToChan(wg *sync.WaitGroup, chanReceive chan model.SimpleData, chanSend chan model.SimpleData) {
	for content := range chanSend {
		chanReceive <- content
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int, chanReceive chan model.SimpleData, chanSend chan model.SimpleData) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go receiveDataToChan(&wg, chanReceive, chanSend)
	}
	wg.Wait()
	close(chanReceive)
}

func saveFile(done chan bool, chanReceive chan model.SimpleData) {
	for content := range chanReceive {
		utils.WriteFile("output.txt", content)
	}
	done <- true
}

func run() {
	var (
		chanSend    = make(chan model.SimpleData)
		chanReceive = make(chan model.SimpleData)
	)

	go sendDataToChan(chanSend, "test.txt")
	done := make(chan bool)
	go saveFile(done, chanReceive)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers, chanReceive, chanSend)
	<-done
}

func main() {
	run()

}
