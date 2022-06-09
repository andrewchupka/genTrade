package integrations

import (
	"context"
	"fmt"
	"reflect"

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


func (finn *Finnhub) Lookup(symb string) {
	res, _, err := finn.client.SymbolSearch(context.Background()).Q(symb).Execute()
	if (err == nil) {
		fmt.Printf("%+v\n", res)
		fmt.Println(reflect.TypeOf(res))
	}
	
}