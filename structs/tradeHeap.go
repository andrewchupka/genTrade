package structs

import (
	"errors"
	// "fmt"
)

type TradeHeap struct {
	tradesList []Trade
	length int
}

func MakeTradeHeap() *TradeHeap {
	tradeHeap := new(TradeHeap)
	tradeHeap.tradesList = make([]Trade, 100) 
	tradeHeap.length = 0
	return tradeHeap

}

func (tradeHeap *TradeHeap) Push(trade *Trade) {
	// expand the size of the array
	if (tradeHeap.length + 1 == len(tradeHeap.tradesList)) {
		newList := make([]Trade, len(tradeHeap.tradesList) * 2)
		copy(newList, tradeHeap.tradesList)
		tradeHeap.tradesList = newList
	}

	tradeHeap.tradesList[tradeHeap.length] = *trade
	tradeHeap.length += 1

	tradeHeap.heapifyUp()

	// for _, item := range tradeHeap.tradesList {  
	// 	fmt.Printf("%s\n", item.Stringify())  
	// }  

}

func (tradeHeap *TradeHeap) Pop() (Trade, error) {
	// check if the trades Heap is empty
	if (tradeHeap.length == 0){
		return *new(Trade), errors.New("TRADES HEAP IS EMPTY")
	}

	item := tradeHeap.tradesList[0]
	tradeHeap.tradesList[0] = tradeHeap.tradesList[tradeHeap.length - 1]
	tradeHeap.length -= 1

	tradeHeap.heapifyDown()

	return item, nil

}

func (tradeHeap *TradeHeap) GetData() []Trade {
	return tradeHeap.tradesList
}

func (tradeHeap *TradeHeap) heapifyDown() {
	index := 0

	for(tradeHeap.hasLeftChild(index)) {
		smallerChildIndex := getLeftChildIndex(index)
		if (tradeHeap.hasRightChild(index) && compareTrades(tradeHeap.getLeftChild(index), tradeHeap.getRightChild(index))) {
			smallerChildIndex = getRightChildIndex(index)
		}

		if (compareTrades(tradeHeap.tradesList[smallerChildIndex], tradeHeap.tradesList[index])) {
			break
		}

		tradeHeap.swap(index, smallerChildIndex)
		index = smallerChildIndex
	}
}

func (tradeHeap *TradeHeap) heapifyUp() {
	index := tradeHeap.length - 1

	for(hasParent(index) && compareTrades(tradeHeap.getParent(index), tradeHeap.tradesList[index])) {
		tradeHeap.swap(getParentIndex(index), index)
		index = getParentIndex(index)
	}
}

func (tradeHeap *TradeHeap) swap(index1 int, index2 int) {
	temp := tradeHeap.tradesList[index1]
	tradeHeap.tradesList[index1] = tradeHeap.tradesList[index2]
	tradeHeap.tradesList[index2] = temp
} 

func hasParent(index int) bool {
	return getParentIndex(index) >= 0
}

func (tradeHeap *TradeHeap) hasLeftChild(index int) bool {
	return getLeftChildIndex(index) < tradeHeap.length
}

func (tradeHeap *TradeHeap) hasRightChild(index int) bool {
	return getRightChildIndex(index) < tradeHeap.length
}

func (tradeHeap *TradeHeap) getParent(index int) Trade {
	return tradeHeap.tradesList[getParentIndex(index)]
}

func (tradeHeap *TradeHeap) getLeftChild(index int) Trade {
	return tradeHeap.tradesList[getLeftChildIndex(index)]
}

func (tradeHeap *TradeHeap) getRightChild(index int) Trade {
	return tradeHeap.tradesList[getRightChildIndex(index)]
}

func getLeftChildIndex(index int) int {
	return (2 * index + 1)
}

func getRightChildIndex(index int) int {
	return (2 * index + 2)
}

func getParentIndex(index int) int {
	return (index - 1) / 2
}

func compareTrades(trade1 Trade, trade2 Trade) bool {
	return trade2.Timestamp.Before(trade1.Timestamp)
}
 
