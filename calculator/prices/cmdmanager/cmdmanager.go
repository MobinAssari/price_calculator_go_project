package cmdmanager

import "fmt"

type CMDManager struct{}

func (cmd CMDManager) WriteResult(data any) error {
	fmt.Print(data)
	return nil
}
func New() CMDManager {
	return CMDManager{}
}
func (cmd CMDManager) ReadFile() ([]string, error) {
	var prices []string
	for {
		var price string
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}
