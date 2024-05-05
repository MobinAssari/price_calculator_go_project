package prices

import (
	"fmt"

	"example.com/calculator/prices/conversion"
	"example.com/calculator/prices/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager        iomanager.IOManager `json:"-"`
	TaxRate          float64             `josn:"tax_rate"`
	InputPrices      []float64           `json:"input_price"`
	TaxIncludedPrice map[string]string   `json:"tax_included_price"`
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadFile()
	prices := make([]float64, len(lines))
	prices, err = conversion.StringToFloat(lines)
	if err != nil {
		println(err)
		return
	}
	job.InputPrices = prices
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		TaxRate:   taxRate,
		//InputPrices: []float64{10, 20, 30},
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		jobTax := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", jobTax)
	}
	job.TaxIncludedPrice = result
	job.IOManager.WriteResult(job)
}
