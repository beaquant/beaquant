package data

import "testing"

func TestDepthEventFromCSVeData_Load(t *testing.T) {
	symbol := "depth_binance.com_BTC_USDT_2020-01-21.csv"
	data := NewDepthEventFromCSVeData("sample-data/")
	t.Log(data.Load(symbol))
	t.Log(data)
}
