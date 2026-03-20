package cmdmanager

import "fmt"

type CmdManager struct {
}

func (cmd CmdManager) ReadLines() ([]string, error) {
	var prices []string
	for {
		var price string
		fmt.Println("price")
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (cmd CmdManager) WriteJSON(data any) error {
	fmt.Println(data)
	return nil
}

func New() CmdManager {
	return CmdManager{}
}
