package main

import (
	"fmt"
	"priceCalculator/project/fileManager"
	"priceCalculator/project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 017}

	for _, taxRate := range taxRates {
		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdM=cmdmanager.New()
		pricejob := prices.NewTaxIncludedPricesjob(fm, taxRate)
		err := pricejob.Process()

		if err != nil {
			fmt.Println("coudnt process")
			fmt.Println(err)
		}

	}
}
