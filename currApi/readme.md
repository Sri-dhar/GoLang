# Currency Exchange TUI

A simple terminal-based application for checking currency exchange rates and performing conversions.

## Features

- Fetch live currency exchange rates from Open Exchange Rates API
- Convert between different currencies
- User-friendly terminal interface

## Installation

1. Ensure you have Go installed on your system (1.16+ recommended)
2. Clone this repository
3. Install dependencies:

```bash
go get github.com/charmbracelet/bubbles/spinner
go get github.com/charmbracelet/bubbles/textinput
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
```

## Configuration

Add your API key to the environment before running:

```bash
export EXCHANGE_API_KEY=your_api_key_here
```

Or modify the code to use environment variables:

```go
apiKey := os.Getenv("EXCHANGE_API_KEY")
```

## Usage

1. Build the executable:

```bash
go build -o currency-exchange
```

2. Run the application:

```bash
./currency-exchange
```

3. Navigate the application:
   - Press `1` to refresh exchange rates
   - Press `2` to convert between currencies
   - Use `Tab` to navigate between input fields
   - Press `Enter` to confirm
   - Press `Esc` to go back to main menu
   - Press `q` to quit
