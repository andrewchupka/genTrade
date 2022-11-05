package structs

import (
	"fmt"
	"genTrade/helpers"
	"math/rand"
	"os"
	"time"
)

type TradeGene struct {
	volumeThreshold float64
	velocityThreshold float64
	geneMatrix [8][8]int
	TradeChannel chan Trade
}

func MakeGene() TradeGene {
	rand.Seed(time.Now().UnixNano())
	tradeGene := new(TradeGene)
	tradeGene.volumeThreshold = generateSizeTreshold()
	tradeGene.velocityThreshold = generateVelocityThreshold()
	tradeGene.geneMatrix = generateGeneMatrix()
	tradeGene.TradeChannel = make(chan Trade, helpers.TRADE_CHANNEL_LEN)
	return *tradeGene
}

func MakeGeneFromFile() *TradeGene {
	return new(TradeGene)
}

func (tradeGene *TradeGene) Process() {
	for {
		message := <-tradeGene.TradeChannel
		fmt.Println("In TradeGene: ", message)
	}
}

func (tradeGene *TradeGene) WriteToFile(file os.File) {

	file.WriteString(fmt.Sprintf("%f\n", tradeGene.volumeThreshold))
	file.WriteString(fmt.Sprintf("%f\n", tradeGene.velocityThreshold))
	
	for _, array := range(tradeGene.geneMatrix) {
		for _, item := range(array) {
			file.WriteString(fmt.Sprintf("%d ", item))
		}
		file.WriteString("\n")
	}
}

func generateSizeTreshold() float64 {
	return rand.ExpFloat64()
}

func generateVelocityThreshold() float64 {
	return rand.ExpFloat64()
}

func generateGeneMatrix() [8][8]int {
	var matrix [8][8]int

	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = rand.Intn(201) - 100
		}
	} 

	return matrix
}