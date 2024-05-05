package main

import (
	"fmt"

	"example.com/calculator/prices"
	"example.com/calculator/prices/filemanager"
)

func main() {

	taxRates := []float64{0.7, 0.9, 1.5}
	channels := make([]chan bool, len(taxRates))

	for index, taxRate := range taxRates {
		channels[index] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%v.txt", taxRate))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(channels[index])
	}
	for _, channel := range channels {
		<-channel
	}
}
