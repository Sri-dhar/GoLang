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

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
)

var API_KEY string

func init() {
	godotenv.Load()
	API_KEY = os.Getenv("EXCHANGE_API_KEY")
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1)

	infoStyle = lipgloss.NewStyle().
			Italic(true).
			Foreground(lipgloss.Color("#666666"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000"))

	highlightStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7D56F4")).
			Bold(true)

	rateStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575"))
)

type ExchangeRates struct {
	Base     string             `json:"base"`
	Rates    map[string]float64 `json:"rates"`
	LastTime time.Time
}

type Page int

const (
	MenuPage Page = iota
	RefreshPage
	ConvertPage
)

type model struct {
	page           Page
	rates          ExchangeRates
	spinner        spinner.Model
	loading        bool
	err            error
	fromInput      textinput.Model
	toInput        textinput.Model
	amountInput    textinput.Model
	activeInput    int
	conversionInfo string
	currencies     []string
	selectedFrom   int
	selectedTo     int
}

type fetchRatesMsg struct {
	rates ExchangeRates
	err   error
}

func fetchRates() tea.Cmd {
	return func() tea.Msg {
		resp, err := http.Get("https://openexchangerates.org/api/latest.json?app_id=" + API_KEY)
		if err != nil {
			return fetchRatesMsg{err: err}
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := ioutil.ReadAll(resp.Body)
			return fetchRatesMsg{err: fmt.Errorf("API Error: %s - %s", resp.Status, string(body))}
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fetchRatesMsg{err: err}
		}

		var rates ExchangeRates
		err = json.Unmarshal(body, &rates)
		if err != nil {
			return fetchRatesMsg{err: err}
		}

		rates.LastTime = time.Now()
		return fetchRatesMsg{rates: rates, err: nil}
	}
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	fromInput := textinput.New()
	fromInput.Placeholder = "USD"
	fromInput.Focus()
	fromInput.Width = 10

	toInput := textinput.New()
	toInput.Placeholder = "EUR"
	toInput.Width = 10

	amountInput := textinput.New()
	amountInput.Placeholder = "1.00"
	amountInput.Width = 20

	return model{
		page:        MenuPage,
		spinner:     s,
		fromInput:   fromInput,
		toInput:     toInput,
		amountInput: amountInput,
		currencies:  []string{"USD", "EUR", "GBP", "JPY", "CAD", "AUD", "CNY", "INR"},
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		fetchRates(),
		m.spinner.Tick,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "1":
			if m.page == MenuPage {
				m.page = RefreshPage
				m.loading = true
				return m, fetchRates()
			}
		case "2":
			if m.page == MenuPage {
				m.page = ConvertPage
				m.fromInput.Focus()
				m.activeInput = 0
				return m, nil
			}
		case "esc":
			if m.page != MenuPage {
				m.page = MenuPage
				return m, nil
			}
		case "tab", "shift+tab", "enter", "up", "down":
			if m.page == ConvertPage {
				if msg.String() == "enter" {
					m.convert()
					return m, nil
				}

				if msg.String() == "up" || msg.String() == "down" {
					return m, nil
				}

				if msg.String() == "tab" || msg.String() == "shift+tab" {
					if msg.String() == "tab" {
						m.activeInput = (m.activeInput + 1) % 3
					} else {
						m.activeInput = (m.activeInput - 1 + 3) % 3
					}

					for i := 0; i < 3; i++ {
						if i == m.activeInput {
							switch i {
							case 0:
								m.fromInput.Focus()
								m.toInput.Blur()
								m.amountInput.Blur()
							case 1:
								m.fromInput.Blur()
								m.toInput.Focus()
								m.amountInput.Blur()
							case 2:
								m.fromInput.Blur()
								m.toInput.Blur()
								m.amountInput.Focus()
							}
						}
					}
				}
			}
		}

	case fetchRatesMsg:
		m.loading = false
		if msg.err != nil {
			m.err = msg.err
		} else {
			m.rates = msg.rates
			m.err = nil
		}
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	if m.page == ConvertPage {
		switch m.activeInput {
		case 0:
			m.fromInput, cmd = m.fromInput.Update(msg)
		case 1:
			m.toInput, cmd = m.toInput.Update(msg)
		case 2:
			m.amountInput, cmd = m.amountInput.Update(msg)
		}
		return m, cmd
	}

	return m, cmd
}

func (m *model) convert() {
	if len(m.rates.Rates) == 0 {
		m.conversionInfo = "No exchange rate data available. Please refresh data first."
		return
	}

	fromCurrency := strings.ToUpper(m.fromInput.Value())
	toCurrency := strings.ToUpper(m.toInput.Value())
	amountStr := m.amountInput.Value()

	if fromCurrency == "" || toCurrency == "" || amountStr == "" {
		m.conversionInfo = "Please fill in all fields"
		return
	}

	if _, exists := m.rates.Rates[fromCurrency]; !exists && fromCurrency != "USD" {
		m.conversionInfo = fmt.Sprintf("Currency %s not found", fromCurrency)
		return
	}

	if _, exists := m.rates.Rates[toCurrency]; !exists && toCurrency != "USD" {
		m.conversionInfo = fmt.Sprintf("Currency %s not found", toCurrency)
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		m.conversionInfo = "Invalid amount"
		return
	}

	var result float64

	if fromCurrency == "USD" {
		result = amount * m.rates.Rates[toCurrency]
	} else if toCurrency == "USD" {
		result = amount / m.rates.Rates[fromCurrency]
	} else {
		usdAmount := amount / m.rates.Rates[fromCurrency]
		result = usdAmount * m.rates.Rates[toCurrency]
	}

	m.conversionInfo = fmt.Sprintf("%.2f %s = %.2f %s",
		amount, fromCurrency, result, toCurrency)
}

func (m model) View() string {
	var s string

	switch m.page {
	case MenuPage:
		s = titleStyle.Render("CURRENCY EXCHANGE") + "\n\n"

		if !m.rates.LastTime.IsZero() {
			s += infoStyle.Render(fmt.Sprintf("Last updated: %s\n\n",
				m.rates.LastTime.Format("Jan 02, 2006 15:04:05")))
			s += infoStyle.Render(fmt.Sprintf("\n"))
		}
		menuStyle := lipgloss.NewStyle().Align(lipgloss.Left)
		s += menuStyle.Render("Choose an option:") + "\n\n"
		s += highlightStyle.Render("1.") + " Refresh Exchange Rates\n"
		s += highlightStyle.Render("2.") + " Convert Currency\n\n"
		s += "Press " + highlightStyle.Render("q") + " to quit\n\n"

		if len(m.rates.Rates) > 0 {
			s += highlightStyle.Render("Available rates (vs USD):\n")
			s += highlightStyle.Render("\n")

			currencies := []string{"EUR", "GBP", "JPY", "CAD", "AUD", "CNY", "INR"}
			for _, curr := range currencies {
				if rate, ok := m.rates.Rates[curr]; ok {
					s += fmt.Sprintf("%s: %s\n",
						curr,
						rateStyle.Render(fmt.Sprintf("%.4f", rate)))
				}
			}
		}

	case RefreshPage:
		s = titleStyle.Render("REFRESHING DATA") + "\n\n"

		if m.loading {
			s += m.spinner.View() + " Fetching latest exchange rates...\n"
		} else {
			if m.err != nil {
				s += errorStyle.Render(fmt.Sprintf("Error: %v\n\n", m.err))
			} else {
				s += "✓ Successfully updated exchange rates!\n\n"
				s += fmt.Sprintf("Base Currency: %s\n\n", m.rates.Base)

				s += "Sample Exchange Rates:\n"
				currencies := []string{"EUR", "GBP", "JPY", "CAD", "AUD", "CNY", "INR"}
				for _, curr := range currencies {
					if rate, ok := m.rates.Rates[curr]; ok {
						s += fmt.Sprintf("%s: %s\n",
							curr,
							rateStyle.Render(fmt.Sprintf("%.4f", rate)))
					}
				}
			}
			s += "\nPress ESC to return to menu"
		}

	case ConvertPage:
		s = titleStyle.Render("CURRENCY CONVERTER") + "\n\n"

		s += fmt.Sprintf("From: %s\n", m.fromInput.View())
		s += fmt.Sprintf("To: %s\n", m.toInput.View())
		s += fmt.Sprintf("Amount: %s\n\n", m.amountInput.View())

		if m.conversionInfo != "" {
			if strings.HasPrefix(m.conversionInfo, "Error") ||
				strings.HasPrefix(m.conversionInfo, "Currency") ||
				strings.HasPrefix(m.conversionInfo, "Invalid") ||
				strings.HasPrefix(m.conversionInfo, "Please") ||
				strings.HasPrefix(m.conversionInfo, "No exchange") {
				s += errorStyle.Render(m.conversionInfo)
			} else {
				s += highlightStyle.Render(m.conversionInfo)
			}
		}

		s += "\n\nPress TAB to switch fields, ENTER to convert, ESC to return to menu"
	}

	if m.err != nil && m.page != RefreshPage {
		s += "\n\n" + errorStyle.Render(fmt.Sprintf("Error: %v", m.err))
	}

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}
