package integrations

import (
	"context"
	"encoding/json"
	"fmt"
	"genTrade/helpers"
	"log"
	"os"
	"os/signal"
	"sync"

	structs "genTrade/structs"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

var SANDBOX_TOKEN string = os.Getenv("FINNHUB_SANDBOX")
var TOKEN string = os.Getenv("FINNHUB_TOKEN")

type Finnhub struct {
	client *finnhub.DefaultApiService
}

func NewFinnhub() *Finnhub {
	config := finnhub.NewConfiguration()
	config.AddDefaultHeader("X-Finnhub-Token", TOKEN)

	finn := new(Finnhub)
	finn.client = finnhub.NewAPIClient(config).DefaultApi
	return finn
}

func (finn *Finnhub) Lookup(symb string) finnhub.SymbolLookup {
	fmt.Println("Performing lookup for ", symb)
	res, _, err := finn.client.SymbolSearch(context.Background()).Q(symb).Execute()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Finished lookup for ", symb)

	return res
}

func (finn *Finnhub) ListLookup(list []string) []finnhub.SymbolLookup {
	var symbolsLength int = len(list)

	var wg sync.WaitGroup
	wg.Add(symbolsLength)

	lookupChannel := make(chan finnhub.SymbolLookup, symbolsLength)

	for _, symbol := range list {

		go func(lookupChannel chan finnhub.SymbolLookup, symbol string, wg *sync.WaitGroup) {
			defer wg.Done()
			lookupChannel <- finn.Lookup(symbol)
		}(lookupChannel, symbol, &wg)

		fmt.Printf("Started lookup for %s\n", symbol)
	}

	wg.Wait()
	close(lookupChannel)

	var resultList []finnhub.SymbolLookup
	for result := range lookupChannel {
		resultList = append(resultList, result)
	}
	return resultList

}

func (finn *Finnhub) BasicFinancials(symb string) finnhub.BasicFinancials {
	fmt.Printf("Performing basic financial lookup for %s\n", symb)

	res, _, err := finn.client.CompanyBasicFinancials(context.Background()).Symbol(symb).Metric("all").Execute()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Finished financial lookup for ", symb)

	return res

}

func (finn *Finnhub) ListBasicFinancials(list []string) []finnhub.BasicFinancials {
	var symbolsLength int = len(list)

	var wg sync.WaitGroup
	wg.Add(symbolsLength)

	lookupChannel := make(chan finnhub.BasicFinancials, symbolsLength)

	for _, symbol := range list {

		go func(lookupChannel chan finnhub.BasicFinancials, symbol string, wg *sync.WaitGroup) {
			defer wg.Done()
			lookupChannel <- finn.BasicFinancials(symbol)
		}(lookupChannel, symbol, &wg)

		fmt.Printf("Started basic financials lookup for %s\n", symbol)
	}

	wg.Wait()
	close(lookupChannel)

	var resultList []finnhub.BasicFinancials
	for result := range lookupChannel {
		resultList = append(resultList, result)
	}
	return resultList

}

func (finn *Finnhub) TradeLookup(symbol string, finishedWriting chan string) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	messages := make(chan string)

	file := helpers.MakeOutputFile(symbol, helpers.TRADE_OUTPUT_DIR)

	trades := structs.MakeTradeHeap()

	defer func(trades structs.TradeHeap, file os.File, finishedWriting chan string) {
		value, _ := json.Marshal(trades.GetData())

		file.WriteString(fmt.Sprintf("%s\n", value))
		file.Close()

		finishedWriting <- file.Name()
	}(*trades, file, finishedWriting)

	go handleWebSocketConnection(symbol, messages, TOKEN)

	for {
		select {
		case <-interrupt:
			log.Println("Interrupted")
			return
		case message := <-messages:
			fmt.Printf("In Finnhub messages: %s\n", message)

			var tradeMessage structs.LiveTrade
			err := json.Unmarshal([]byte(message), &tradeMessage)
			if err != nil {
				log.Printf("Could not unmarshall trade message: {%s}\n", err)
				continue
			}

			if tradeMessage.Type == "trade" {
				for _, item := range tradeMessage.Data {
					trades.Push(&item)
				}
			}
		}
	}
}
