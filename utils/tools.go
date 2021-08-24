package utils

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"Goroutine/model"
)

func ReadFile(filePath string) string {
	randomNum := 3 + rand.Intn(6-3+1)
	log.Println("time read input:", randomNum)
	time.Sleep(time.Duration(randomNum) * time.Second)

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatalf("File reading error", err)
	}
	// defer wg.Done()
	return string(data)
}

func SplitString(data string) model.SimpleData {
	listData := strings.Fields(data)

	randomNum := 6 + rand.Intn(15-6+1)
	log.Println("time handling:", randomNum)
	time.Sleep(time.Duration(randomNum) * time.Second)

	// defer wg.Done()

	return model.SimpleData{listData}
}

func WriteFile(fileName string, content model.SimpleData, wg *sync.WaitGroup) {

	randomNum := 3 + rand.Intn(6-3+1)
	log.Println("time write output:", randomNum)
	time.Sleep(time.Duration(randomNum) * time.Second)

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Fatalln("Error open file", err)
	}
	defer f.Close()

	for i := range content.Data {
		if _, err := f.WriteString(content.Data[i] + "\n"); err != nil {
			log.Println("Error write:", err)
			f.Close()
		}
	}
	err = f.Close()
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully")

	wg.Done()
}
