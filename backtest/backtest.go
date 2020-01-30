package backtest

type Backtest struct {
	exchanges []string
	symbols   []string
	data      []DataHandler
}
