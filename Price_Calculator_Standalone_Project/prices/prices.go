package prices

import (
	"fmt"
	"priceCalculator/project/converter"
	iomanager "priceCalculator/project/ioManager"
)

type TaxIncludedPricesjob struct {
	IOmanager         iomanager.IoManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPricesjob) LoadData() error {

	lines, err := job.IOmanager.ReadLines()
	if err != nil {
		// fmt.Println(err)

		return err
	}

	prices, err := converter.StringToFloats(lines)

	if err != nil {
		// fmt.Println(err)

		return err
	}

	job.InputPrices = prices

	return  nil

}

func (job *TaxIncludedPricesjob) Process() error {
	err :=job.LoadData()
	if err != nil{
		return err
	}
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	// fmt.Println(result)
	job.TaxIncludedPrices = result
	return job.IOmanager.WriteJSON(job)
	// fmt.Println("Saved")
}

func NewTaxIncludedPricesjob(fm iomanager.IoManager, taxRate float64) *TaxIncludedPricesjob {
	return &TaxIncludedPricesjob{
		IOmanager: fm,
		TaxRate:   taxRate,
	}
}
