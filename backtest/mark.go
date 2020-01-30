package backtest

import (
	"time"
)

// MarkHandler declares the basic event interface
type MarkHandler interface {
	Time() time.Time
	SetTime(time.Time)
	Symbol() string
	SetSymbol(string)
	Exchange() string
	SetExchange(string)
}

// Mark is the implementation of the basic event interface.
type Mark struct {
	timestamp time.Time
	symbol    string
	exchange  string
}

// Time returns the timestamp of an event
func (m Mark) Time() time.Time {
	return m.timestamp
}

// SetTime returns the timestamp of an event
func (m *Mark) SetTime(t time.Time) {
	m.timestamp = t
}

// Symbol returns the symbol string of the event
func (m Mark) Symbol() string {
	return m.symbol
}

// SetSymbol returns the symbol string of the event
func (m *Mark) SetSymbol(s string) {
	m.symbol = s
}

// Symbol returns the symbol string of the event
func (m Mark) Exchange() string {
	return m.exchange
}

// SetSymbol returns the symbol string of the event
func (m *Mark) SetExchange(e string) {
	m.exchange = e
}
