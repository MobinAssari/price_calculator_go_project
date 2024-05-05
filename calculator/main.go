package main

import (
	"fmt"

	"example.com/calculator/prices"
	"example.com/calculator/prices/filemanager"
)

func main() {

	taxRates := []float64{0.7, 0.9, 1.5}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%v.txt", taxRate))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}
}
