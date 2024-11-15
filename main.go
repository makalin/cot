package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

const binanceAPIURL = "https://api.binance.com/api/v3/ticker/price?symbol="

type Coin struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

type Portfolio struct {
	Coins     []Coin             `json:"coins"`
	Timestamp time.Time          `json:"timestamp"`
	Alarms    map[string]float64 `json:"alarms"`
}

func getCoinPrice(symbol string) (float64, error) {
	resp, err := http.Get(binanceAPIURL + symbol)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var response map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, err
	}

	price, err := strconv.ParseFloat(response["price"], 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func loadPortfolio(filename string) (Portfolio, error) {
	var portfolio Portfolio
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return portfolio, nil 
		}
		return portfolio, err
	}

	if err := json.Unmarshal(file, &portfolio); err != nil {
		return portfolio, err
	}

	return portfolio, nil
}

func savePortfolio(filename string, portfolio Portfolio) error {
	portfolio.Timestamp = time.Now()
	file, err := json.MarshalIndent(portfolio, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, file, 0644)
}

func addCoin(portfolio *Portfolio, symbol string) {
	for _, coin := range portfolio.Coins {
		if coin.Symbol == symbol {
			fmt.Println("Coin already exists in portfolio.")
			return
		}
	}
	portfolio.Coins = append(portfolio.Coins, Coin{Symbol: symbol})
	fmt.Println("Coin added successfully.")
}

func listCoins(portfolio *Portfolio) {
	if len(portfolio.Coins) == 0 {
		fmt.Println("No coins in portfolio.")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Symbol", "Current Price"})

	for _, coin := range portfolio.Coins {
		price, err := getCoinPrice(coin.Symbol)
		if err != nil {
			fmt.Printf("Error fetching price for %s: %v\n", coin.Symbol, err)
			continue
		}
		row := []string{coin.Symbol, fmt.Sprintf("$%.2f", price)}
		table.Append(row)
	}

	table.Render()
}

func removeCoin(portfolio *Portfolio, symbol string) {
	for i, coin := range portfolio.Coins {
		if coin.Symbol == symbol {
			portfolio.Coins = append(portfolio.Coins[:i], portfolio.Coins[i+1:]...)
			fmt.Println("Coin removed.")
			return
		}
	}
	fmt.Println("Coin not found in the portfolio.")
}

func setAlarm(portfolio *Portfolio, symbol string, price float64) {
	if portfolio.Alarms == nil {
		portfolio.Alarms = make(map[string]float64)
	}
	portfolio.Alarms[symbol] = price
	fmt.Printf("Alarm set for %s at $%.2f\n", symbol, price)
}

func checkAlarms(portfolio *Portfolio) {
	for symbol, alarmPrice := range portfolio.Alarms {
		currentPrice, err := getCoinPrice(symbol)
		if err != nil {
			fmt.Printf("Error fetching price for %s: %v\n", symbol, err)
			continue
		}
		if currentPrice >= alarmPrice {
			fmt.Printf("ALARM: %s has reached $%.2f (current: $%.2f)\n", symbol, alarmPrice, currentPrice)
		}
	}
}

func displayHelp() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Command", "Description"})

	data := [][]string{
		{"add <symbol>", "Add a coin to your portfolio (e.g., BTCUSDT)."},
		{"list", "List coins in portfolio."},
		{"remove <symbol>", "Remove coin from portfolio."},
		{"alarm <symbol> <price>", "Set a price alarm for a coin."},
		{"save", "Save portfolio config file."},
		{"alarms", "Check all set alarms."},
		{"help", "Display this help."},
	}

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
}

func main() {
	portfolio, err := loadPortfolio("portfolio.json")
	if err != nil {
		fmt.Println("Error loading portfolio:", err)
		return
	}

	if len(os.Args) < 2 {
		displayHelp()
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) != 3 {
			fmt.Println("Usage: add <symbol>")
			return
		}
		symbol := strings.ToUpper(os.Args[2])
		addCoin(&portfolio, symbol)
	case "list":
		listCoins(&portfolio)
	case "remove":
		if len(os.Args) != 3 {
			fmt.Println("Usage: remove <symbol>")
			return
		}
		symbol := strings.ToUpper(os.Args[2])
		removeCoin(&portfolio, symbol)
	case "alarm":
		if len(os.Args) != 4 {
			fmt.Println("Usage: alarm <symbol> <price>")
			return
		}
		symbol := strings.ToUpper(os.Args[2])
		price, err := strconv.ParseFloat(os.Args[3], 64)
		if err != nil {
			fmt.Println("Invalid price. Please enter a valid number.")
			return
		}
		setAlarm(&portfolio, symbol, price)
	case "save":
		if err := savePortfolio("portfolio.json", portfolio); err != nil {
			fmt.Println("Error saving portfolio:", err)
		} else {
			fmt.Println("Portfolio saved successfully.")
		}
	case "alarms":
		checkAlarms(&portfolio)
	case "help":
		displayHelp()
	default:
		fmt.Println("Unknown command. Type 'help' for a list of commands.")
	}
}
