package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"
)

type Output struct {
	Date           string  `json:"date"`
	BaseCurrency   string  `json:"base_currency"`
	TargetCurrency string  `json:"target_currency"`
	ExchangeRate   float64 `json:"exchange_rate"`
}

func main() {
	var input, output string
	if len(os.Args) < 3 {
		input = "input.json"
		output = "output.json"
	} else {
		input = os.Args[1]
		output = os.Args[2]
	}

	clean(input, output)
	curate(input, output)
}

func clean(input, output string) {
	inputFile, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(output)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	outputData := Output{"", "EUR", "", 0.0}
	decoder := json.NewDecoder(inputFile)
	encoder := json.NewEncoder(outputFile)

	for decoder.More() {
		var input map[string]map[string]float64
		err := decoder.Decode(&input)
		if err != nil {
			fmt.Println(err)
			return
		}

		for date, currencies := range input {
			outputData.Date = date
			for currency, exchangeRate := range currencies {
				outputData.TargetCurrency = currency
				outputData.ExchangeRate = exchangeRate
				err := encoder.Encode(outputData)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}

func curate(input, output string) {
	inputFile, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(output)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	outputData := make(map[string][]Output)
	decoder := json.NewDecoder(inputFile)
	encoder := json.NewEncoder(outputFile)

	for decoder.More() {
		var input []Output
		err := decoder.Decode(&input)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, data := range input {
			if _, ok := outputData[data.TargetCurrency]; !ok {
				outputData[data.TargetCurrency] = []Output{}
			}
			outputData[data.TargetCurrency] = append(outputData[data.TargetCurrency], data)
		}
	}

	for _, data := range outputData {
		sort.Slice(data, func(i, j int) bool {
			return data[i].Date < data[j].Date
		})

		lastExchangeRate := data[0].ExchangeRate
		lastDate, _ := time.Parse("2006-01-02", data[0].Date)

		for index, _ := range data {
			if !nextDateExist(lastDate, data) {
				output := Output{"", "EUR", "", 0.0}

				output.Date = lastDate.AddDate(0, 0, 1).Format("2006-01-02")
				output.TargetCurrency = data[index].TargetCurrency
				output.ExchangeRate = lastExchangeRate

				data = append(data, output)
			}

			lastDate, _ = time.Parse("2006-01-02", data[index].Date)
			lastExchangeRate = data[index].ExchangeRate
		}

		for _, data := range data {
			err := encoder.Encode(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func nextDateExist(date time.Time, data []Output) bool {
	nextDate := date.AddDate(0, 0, 1).Format("2006-01-02")
	for _, d := range data {
		if d.Date == nextDate {
			return true
		}
	}
	return false
}
