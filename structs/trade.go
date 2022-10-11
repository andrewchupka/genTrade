package structs

import (
	"time"
	"encoding/json"
	"log"
)

type Trade struct {
	Price float64 `json:p`
	Timestamp time.Time `json:t`
	Symbol string `json:s`
	Volume float64 `json: v`
	additionalData interface{} `json:c`
}

func (trade *Trade) Stringify() string {
	tradeJson, err := json.Marshal(trade)
	if (err != nil) {
		log.Fatal(err.Error())
	} 
	
	return string(tradeJson)
}

