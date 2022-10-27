package structs

import (
	"fmt"
	"log"
	"os"
	"genTrade/helpers"
)

type GenePool struct {
	pool [helpers.POOL_POPULATION]TradeGene
	symbol string
}

func MakeGenePool(symbol string) *GenePool {
	genePool := new(GenePool)
	genePool.symbol = symbol
	for i := 0; i < helpers.POOL_POPULATION; i++ {
		genePool.pool[i] = MakeGene()
	}

	return genePool
}

func (genePool *GenePool) ExportGenePool() {

	outputDir := helpers.PathSanitizer(fmt.Sprintf("%s%s/", helpers.GENE_POOL_DIR, genePool.symbol))

	fmt.Println("Removing existing gene data")
	err := os.RemoveAll(outputDir)

	if err != nil {
		log.Fatal("Unable to clear the gene pool")
		panic(err)
	}

	err = os.Mkdir(outputDir, 0755)
	if err != nil {
		log.Fatal("Unable to remake the gene pool directory")
		panic(err)
	}

	fmt.Println("Removed existing gene data")

	fmt.Println("Generating genetic data")
	for i, item := range genePool.pool {
		file := helpers.MakeOutputFile(fmt.Sprintf("%d", i), fmt.Sprintf("%s%s/", helpers.GENE_POOL_DIR, genePool.symbol))
		item.WriteToFile(file)
		fmt.Printf("Finished writing %s\n", file.Name())
		file.Close()
	}
	fmt.Println("Finished writing genetic data")

}
