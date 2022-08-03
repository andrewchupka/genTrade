package integrations

import (
	"context"
	"fmt"
	"sync"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

const TOKEN string = "sandbox_cagtomiad3i02fchd2ng"

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
	if (err != nil) {
		fmt.Println(err)
		panic(err)
	}
	
	fmt.Println("Finished lookup for ", symb)

	return res
}

func (finn *Finnhub) ListLookup(list []string) []finnhub.SymbolLookup{
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
	for  result := range lookupChannel {
		resultList = append(resultList, result)
	}
	return resultList

}

func (finn *Finnhub) BasicFinancials(symb string) finnhub.BasicFinancials{
	fmt.Printf("Performing basic financial lookup for %s\n", symb)

	res, _, err := finn.client.CompanyBasicFinancials(context.Background()).Symbol(symb).Metric("all").Execute()
	if (err != nil) {
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
	for  result := range lookupChannel {
		resultList = append(resultList, result)
	}
	return resultList
	
}

func (finn *Finnhub) TradeLookup(symbol string) {
	messages := make(chan string)

	handleWebSocketConnection(symbol, messages, TOKEN)

	for {
		fmt.Printf("In Finnhub messages: %s", <-messages )
	}
}
