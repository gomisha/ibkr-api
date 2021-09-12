package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Query struct {
	symbol string
	strike string
	month  string
}

func main() {
	query := Query{"FB", "380.0", "OCT21"}

	fmt.Print("query=")
	fmt.Printf("%+v\n", query) //https://pkg.go.dev/fmt

	getConIDBySymbol("SWBI")
	// marketData := getMarketData()
	// fmt.Println("marketData.Bid=" + marketData[0].Bid)
	// fmt.Println("marketData.Ask=" + marketData[0].Ask)
	// fmt.Println("marketData.Description=" + marketData[0].Description)
}

type Symbol struct {
	Symbol string `json:"symbol"`
}

func getMarketData() MarketData {
	//https://stackoverflow.com/a/12122718/5719544
	//disable SSL check - running local IBKR REST gateway
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("https://localhost:5000/v1/api/iserver/marketdata/snapshot?conids=507958489&fields=31,86,7220")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to struct
	var marketData MarketData
	json.Unmarshal(bodyBytes, &marketData)
	fmt.Printf("API Response as struct %+v\n", marketData)

	return marketData
}

//POST REST call
func getConIDBySymbol(symbol string) {
	fmt.Println("2. Performing Http Post...")
	//todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	symbol1 := Symbol{symbol}
	//jsonReq, err := json.Marshal(todo)
	jsonReq, err := json.Marshal(symbol1)
	//resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Post("https://localhost:5000/v1/api/iserver/secdef/search", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert response body to Todo struct
	//var todoStruct Todo
	var searchBySymbol SearchBySymbol
	//json.Unmarshal(bodyBytes, &todoStruct)
	json.Unmarshal(bodyBytes, &searchBySymbol)
	//fmt.Printf("%+v\n", todoStruct)
	fmt.Printf("%+v\n", searchBySymbol)

	fmt.Println("ConID=" + strconv.Itoa(searchBySymbol[0].Conid))
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
//auto generated by https://mholt.github.io/json-to-go/
type MarketData []struct {
	Last        string  `json:"31"`
	Bid         string  `json:"84"`
	AskSize     string  `json:"85"`
	Ask         string  `json:"86"`
	Volume      string  `json:"87"`
	Num6119     string  `json:"6119"`
	Num6508     string  `json:"6508"`
	Num6509     string  `json:"6509"`
	Description string  `json:"7220"`
	ConidEx     string  `json:"conidEx"`
	Conid       int     `json:"conid"`
	Updated     int64   `json:"_updated"`
	ServerID    string  `json:"server_id"`
	Eight7Raw   float64 `json:"87_raw"`
}

type SearchBySymbol []struct {
	Conid         int         `json:"conid"`
	CompanyHeader string      `json:"companyHeader"`
	CompanyName   string      `json:"companyName"`
	Symbol        string      `json:"symbol"`
	Description   string      `json:"description"`
	Restricted    interface{} `json:"restricted"`
	Fop           interface{} `json:"fop"`
	Opt           string      `json:"opt"`
	War           string      `json:"war"`
	Sections      []struct {
		SecType  string `json:"secType"`
		Months   string `json:"months,omitempty"`
		Exchange string `json:"exchange,omitempty"`
		Conid    int    `json:"conid,omitempty"`
	} `json:"sections"`
}
