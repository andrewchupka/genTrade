package integrations

import (
	"context"
	"fmt"
	"sync"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

type Finnhub struct {
	client *finnhub.DefaultApiService
}

func NewFinnhub() *Finnhub {
	config := finnhub.NewConfiguration()
	config.AddDefaultHeader("X-Finnhub-Token", "sandbox_cagtomiad3i02fchd2ng")

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
	defer close(lookupChannel)

	for _, symbol := range list {

		go func(lookupChannel chan finnhub.SymbolLookup, symbol string, wg *sync.WaitGroup) {
			defer wg.Done()
			lookupChannel <- finn.Lookup(symbol)
		}(lookupChannel, symbol, &wg)

		fmt.Printf("Started lookup for %s\n", symbol)
	}

	wg.Wait()

	var resultList []finnhub.SymbolLookup
	for  result := range lookupChannel {
		resultList = append(resultList, result)
	}
	return resultList

}