package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var wp sync.WaitGroup

func responseSize(url string) {
	defer wp.Done()
	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)
	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: ", len(body))
}

func main() {
	wp.Add(1)
	go responseSize("https://www.duoke360.com")
	wp.Add(1)
	go responseSize("https://baidu.com")
	wp.Add(1)
	go responseSize("https://jd.com")

	wp.Wait()
	//time.Sleep(10 * time.Second)
}