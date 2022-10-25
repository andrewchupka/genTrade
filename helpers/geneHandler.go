package helpers

import (
	"os"
	"log"
	"fmt"
	"genTrade/structs"
)

func GenerateIntialPool(genes *[]structs.TradeGene) {

	fmt.Println("Removing existing gene data")
	err := os.RemoveAll(GENE_POOL_DIR)

	if err != nil {
		log.Fatal("Unable to clear the gene pool")
		panic(err)
	}

	err = os.Mkdir(GENE_POOL_DIR, 0755)
	if err != nil {
		log.Fatal("Unable to remake the gene pool directory")
		panic(err)
	}

	fmt.Println("Removed existing gene data")

	fmt.Println("Generating genetic data")
	for i, item := range(*genes) {
		symbol := CRYPTO_LIST[i % len(CRYPTO_LIST)]
		file := MakeOutputFile(fmt.Sprintf("%s_%d", symbol, i), GENE_POOL_DIR)
		item.WriteToFile(file)
		fmt.Printf("Finished writing %s\n", file.Name())
		file.Close()
	}
	fmt.Println("Finished writing genetic data")
	
}

