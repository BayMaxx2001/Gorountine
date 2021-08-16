package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gorountines/model"
)

func ReadFile(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("File reading error", err)
	}

	return string(data)
}

func SplitString(data string) model.SimpleData {
	listData := strings.Fields(data)
	return model.SimpleData{listData}
}

func WriteFile(fileName string, content model.SimpleData) {

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
}
