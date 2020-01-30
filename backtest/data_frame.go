package backtest

import "github.com/nntaoli-project/GoEx"

// DataFrame declares a data event interface
type DataFrame interface {
	MarkHandler
}

// DepthFrame declares a depth stream interface.
type DepthFrame interface {
	DataFrame
}

// Depth declares a data stream for a depth.
type Depth struct {
	Mark
	goex.Depth
}

// Price returns the middle of Bid and Ask.
func (t Depth) Price() float64 {
	latest := (t.BidList[0].Price + t.AskList[0].Price) / float64(2)
	return latest
}

// Spread returns the difference or spread of Bid and Ask.
func (t Depth) Spread() float64 {
	return t.AskList[0].Price - t.BidList[0].Price
}
