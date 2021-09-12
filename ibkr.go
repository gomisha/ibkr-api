package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	get()
	//post()
}

func get() {
	//https://stackoverflow.com/a/12122718/5719544
	//disable SSL check - running local IBKR REST gateway
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// customTransport := http.DefaultTransport.(*http.Transport).Clone()
	// customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// client = &http.Client{Transport: customTransport}

	fmt.Println("1. Performing Http Get...")
	//resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	resp, err := http.Get("https://localhost:5000/v1/api/iserver/marketdata/snapshot?conids=507958489&fields=31,86,7220")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	//var todoStruct Todo
	var marketData MarketData
	//json.Unmarshal(bodyBytes, &todoStruct)
	json.Unmarshal(bodyBytes, &marketData)
	//fmt.Printf("API Response as struct %+v\n", todoStruct)
	fmt.Printf("API Response as struct %+v\n", marketData)
}

func post() {
	fmt.Println("2. Performing Http Post...")
	todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	jsonReq, err := json.Marshal(todo)
	resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("%+v\n", todoStruct)
}

// Todo struct
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

//Market Data Snapshot
//https://www.interactivebrokers.com/api/doc.html#tag/Market-Data/paths/~1iserver~1marketdata~1snapshot/get
type MarketData []struct {
	Num31     string  `json:"31"`
	Num84     string  `json:"84"`
	Num85     string  `json:"85"`
	Num86     string  `json:"86"`
	Num87     string  `json:"87"`
	Num6119   string  `json:"6119"`
	Num6508   string  `json:"6508"`
	Num6509   string  `json:"6509"`
	Num7220   string  `json:"7220"`
	ConidEx   string  `json:"conidEx"`
	Conid     int     `json:"conid"`
	Updated   int64   `json:"_updated"`
	ServerID  string  `json:"server_id"`
	Eight7Raw float64 `json:"87_raw"`
}
