package main
import (
    "fmt"
    "time"
    "sync"

    "Goroutine/model"
    "Goroutine/utils"
)

func sendDataToChan(chanSend chan<- model.SimpleData, filePath string) {
    var (
        listData model.SimpleData
        wg sync.WaitGroup
    )
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go func(){
            data := utils.ReadFile(filePath)   
            listData = utils.SplitString(data, &wg)
        }()
    }
    wg.Wait()
    chanSend <- listData
    close(chanSend)
}

func saveFile(done chan bool, chanReceive <-chan model.SimpleData) {

    for content := range chanReceive {
        utils.WriteFile("output.txt", content)
    }
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

