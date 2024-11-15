# Cryptocurrency Portfolio Tracker

![Project Logo](logo.jpg)

A simple command-line cryptocurrency portfolio tracker that retrieves current coin prices, manages your portfolio, sets alarms, and more using the Binance API.

## Description

This project allows users to manage a cryptocurrency portfolio from the command line. You can add, list, and remove coins, set price alarms, and save/load your portfolio to/from a JSON file. The application uses the Binance API to fetch the current prices for specified symbols.

## Features

- **Add Coins**: Add new coins to your portfolio.
- **List Coins**: Display all coins in your portfolio along with their current market price.
- **Remove Coins**: Remove coins from your portfolio.
- **Set Alarms**: Set a price alarm for specific coins.
- **Check Alarms**: Check if any coins in your portfolio have met their alarm conditions.
- **Save/Load Portfolio**: Save your portfolio data to a file and load it on startup.
- **Command-line Interface**: Simple CLI for managing your portfolio.

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/username/cot.git
   cd cot
   ```

2. Build the project:
   ```bash
   go build
   ```

3. Run the application:
   ```bash
   ./cot
   ```

## Usage

### Commands
- `add <symbol>`: Add a coin to your portfolio (e.g., `add BTCUSDT`).
- `list`: Display all coins in the portfolio.
- `remove <symbol>`: Remove a coin from the portfolio (e.g., `remove BTCUSDT`).
- `alarm <symbol> <price>`: Set a price alarm for a coin (e.g., `alarm BTCUSDT 30000`).
- `save`: Save the current state of the portfolio.
- `alarms`: Check all active alarms.
- `help`: Display available commands.

### Example

To add Bitcoin to your portfolio:
```bash
./cot add BTCUSDT
```

To list all coins in your portfolio:
```bash
./cotlist
```

## Configuration

The portfolio data is saved in a JSON file named `portfolio.json` in the same directory as the executable.

## Dependencies

- [Go](https://golang.org/)
- [Binance API](https://binance.com)
- [olekukonko/tablewriter](https://github.com/olekukonko/tablewriter) - A library for rendering tables in the terminal.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
